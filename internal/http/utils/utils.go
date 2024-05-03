package http_utils

import (
	"bytes"
	"io"
)

func ReadUntilCRLF(r *bytes.Reader) ([]byte, error) {
	return ReadUntil(r, []byte("\r\n"))
}

func ReadCRLF(reader *bytes.Reader, errorMessage string) (err error) {
	offsetBeforeCRLF := GetReaderOffset(reader)

	b, err := reader.ReadByte()
	if err == io.EOF {
		return IncompleteInputError{
			Reader:  reader,
			Message: errorMessage,
		}
	}

	if b != '\r' {
		reader.Seek(int64(offsetBeforeCRLF), io.SeekStart)

		return BadProtocolError{
			Reader:  reader,
			Message: errorMessage,
		}
	}

	b, err = reader.ReadByte()
	if err == io.EOF {
		return IncompleteInputError{
			Reader:  reader,
			Message: errorMessage,
		}
	}

	if b != '\n' {
		reader.Seek(int64(offsetBeforeCRLF), io.SeekStart)

		return BadProtocolError{
			Reader:  reader,
			Message: errorMessage,
		}
	}

	return nil
}

func ReadUntil(r *bytes.Reader, delim []byte) ([]byte, error) {
	var result []byte

	for {
		b, err := r.ReadByte()
		if err != nil {
			if err != io.EOF {
				panic("expected error to always be io.EOF")
			}

			return result, io.EOF
		}

		result = append(result, b)

		if len(result) >= len(delim) && bytes.Equal(result[len(result)-len(delim):], delim) {
			return result[:len(result)-len(delim)], nil
		}
	}
}

func ReadUntilAnyDelimiter(r *bytes.Reader, delimiters [][]byte) ([]byte, error) {
	var result []byte

	for {
		b, err := r.ReadByte()
		if err != nil {
			if err != io.EOF {
				panic("expected error to always be io.EOF")
			}

			return result, io.EOF
		}

		result = append(result, b)

		for _, delim := range delimiters {
			if len(result) >= len(delim) && bytes.Equal(result[len(result)-len(delim):], delim) {
				return result[:len(result)-len(delim)], nil
			}
		}
	}
}

func ReadBytes(r *bytes.Reader, n int) ([]byte, error) {
	var result []byte

	for i := 0; i < n; i++ {
		b, err := r.ReadByte()
		if err != nil {
			if err != io.EOF {
				panic("expected error to always be io.EOF")
			}

			return result, io.EOF
		}

		result = append(result, b)
	}

	return result, nil
}

func ReplaceCharsWithSpace(data []byte, chars [][]byte) []byte {
	var newData []byte
	for _, char := range chars {
		newData = bytes.ReplaceAll(data, char, []byte{' '})
	}
	return newData
}

// for i := 0; i < len(data); i++ {
// 	for _, char := range chars {
// 		if data[i] == char {
// 			data[i] = ' '
// 		}
// 	}
// }
