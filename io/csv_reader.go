package io

import (
	"bufio"
	"context"
	"encoding/csv"
	"errors"
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

	for err == nil {
		select {
		case <-ctx.Done():
			break
		default:
			line, err = reader.Read()
			if err != nil {
				if !errors.Is(io.EOF, err) {
					log.Printf("error reading the csv file: %s", err.Error())
				}
				break
			}
			positionDTO := r.parseLine(line)
			if positionDTO == nil {
				continue
			}

			r.demuxer.Demux(*positionDTO)
		}
	}
	r.demuxer.close()

	return err
}

func (r CSVReader) parseLine(line []string) *PositionDTO {
	rideID, err := strconv.Atoi(line[indexID])
	if err != nil {
		log.Printf("cannot parse rideID: %s. %s", line[indexID], err.Error())
		return nil
	}
	lat, err := strconv.ParseFloat(line[indexLat], 64)
	if err != nil {
		log.Printf("cannot parse Lat: %s. %s", line[indexLat], err.Error())
		return nil
	}
	lon, err := strconv.ParseFloat(line[indexLon], 64)
	if err != nil {
		log.Printf("cannot parse Lon: %s. %s", line[indexLon], err.Error())
		return nil
	}
	timestamp, err := strconv.ParseInt(line[indexTimestamp], 10, 64)
	if err != nil {
		log.Printf("cannot parse timestamp: %s. %s", line[indexTimestamp], err.Error())
		return nil
	}

	return &PositionDTO{
		RideID:    rideID,
		Lat:       lat,
		Lon:       lon,
		Timestamp: timestamp,
	}
}
