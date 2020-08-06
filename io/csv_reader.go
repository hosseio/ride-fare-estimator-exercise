package io

import (
	"bufio"
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const (
	indexID        = 0
	indexLat       = 1
	indexLon       = 2
	indexTimestamp = 3
)

type PositionDTO struct {
	RideID    int
	Lat       float64
	Lon       float64
	Timestamp int64
}

type CSVReader struct {
	demuxer  *Demuxer
	filepath string
}

type CSVFilepath string

func NewCSVReader(demuxer *Demuxer, filepath CSVFilepath) CSVReader {
	return CSVReader{demuxer: demuxer, filepath: string(filepath)}
}

func (r CSVReader) Read(ctx context.Context) error {
	csvFile, err := os.Open(r.filepath)
	if err != nil {
		return err
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var (
		line []string
	)

	defer r.demuxer.close()

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			line, err = reader.Read()
			if err != nil {
				if !errors.Is(io.EOF, err) {
					log.Printf("error reading the csv file: %s", err.Error())
				}
				return nil
			}
			positionDTO, err := r.parseLine(line)
			if err != nil {
				log.Println(err.Error())
				continue
			}
			r.demuxer.Demux(positionDTO)
		}
	}
}

func (r CSVReader) parseLine(line []string) (PositionDTO, error) {
	rideID, err := strconv.Atoi(line[indexID])
	if err != nil {
		return PositionDTO{}, fmt.Errorf("cannot parse rideID: %s. %w", line[indexID], err)
	}

	lat, err := strconv.ParseFloat(line[indexLat], 64)
	if err != nil {
		return PositionDTO{}, fmt.Errorf("cannot parse Lat: %s. %w", line[indexLat], err)
	}

	lon, err := strconv.ParseFloat(line[indexLon], 64)
	if err != nil {
		return PositionDTO{}, fmt.Errorf("cannot parse Lon: %s. %w", line[indexLon], err)
	}

	timestamp, err := strconv.ParseInt(line[indexTimestamp], 10, 64)
	if err != nil {
		return PositionDTO{}, fmt.Errorf("cannot parse timestamp: %s. %w", line[indexTimestamp], err)
	}

	return PositionDTO{
		RideID:    rideID,
		Lat:       lat,
		Lon:       lon,
		Timestamp: timestamp,
	}, nil
}
