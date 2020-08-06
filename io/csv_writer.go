package io

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/hosseio/ride-fare-estimator-exercise/internal/reader"
)

type CSVWriter struct {
	fareRetriever reader.FareRetriever
	filepath      string
}

type CSVOutFilepath string

func NewCSVWriter(fareRetriever reader.FareRetriever, filepath CSVOutFilepath) CSVWriter {
	return CSVWriter{fareRetriever: fareRetriever, filepath: string(filepath)}
}

func (w CSVWriter) Write(ctx context.Context) error {
	csvFile, err := os.Create(w.filepath)
	if err != nil {
		return err
	}
	writer := csv.NewWriter(bufio.NewWriter(csvFile))

	fares, err := w.fareRetriever.Get()
	if err != nil {
		return err
	}
	for _, fare := range fares {
		line := []string{strconv.Itoa(fare.RideID), fmt.Sprintf("%f", fare.Amount)}
		err := writer.Write(line)
		if err != nil {
			log.Printf("error writing in the output: %s", err.Error())
		}
	}

	writer.Flush()

	return nil
}
