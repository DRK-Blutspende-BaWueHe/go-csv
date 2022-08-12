package gocsv

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func encodeMessage(messageData []byte, enc Encoding) ([]byte, error) {
	var messageBytes []byte
	var err error

	switch enc {
	case EncodingUTF8:
		return messageData, nil
	case EncodingASCII:
		return messageData, nil
	case EncodingDOS866:
		messageBytes, err = EncodeCharsetToUTF8From(charmap.CodePage866, messageData)
		return messageBytes, err
	case EncodingDOS855:
		messageBytes, err = EncodeCharsetToUTF8From(charmap.CodePage855, messageData)
		return messageBytes, err
	case EncodingDOS852:
		messageBytes, err = EncodeCharsetToUTF8From(charmap.CodePage852, messageData)
		return messageBytes, err
	case EncodingWindows1250:
		messageBytes, err = EncodeCharsetToUTF8From(charmap.Windows1250, messageData)
		return messageBytes, err
	case EncodingWindows1251:
		messageBytes, err = EncodeCharsetToUTF8From(charmap.Windows1251, messageData)
		return messageBytes, err
	case EncodingWindows1252:
		messageBytes, err = EncodeCharsetToUTF8From(charmap.Windows1252, messageData)
		return messageBytes, err
	case EncodingISO8859_1:
		messageBytes, err = EncodeCharsetToUTF8From(charmap.ISO8859_1, messageData)
		return messageBytes, err
	}

	return []byte{}, fmt.Errorf("invalid Codepage Id='%d' - %w", enc, err)
}

func Unmarshal(messageData []byte, separator string, targetStruct interface{}, enc Encoding, tz Timezone) error {
	var (
		messageBytes []byte
		err          error
	)

	if messageBytes, err = encodeMessage(messageData, enc); err != nil {
		return err
	}

	// identify linebreaks
	bufferedInputLinesWithBlanks := strings.Split(string(messageBytes), string([]byte{0x0A}))
	if len(bufferedInputLinesWithBlanks) <= 1 {
		bufferedInputLinesWithBlanks = strings.Split(string(messageBytes), string([]byte{0x0D}))
	}
	bufferedInputLines := make([]string, 0)
	for _, x := range bufferedInputLinesWithBlanks {
		trimmed := strings.Trim(x, " ")
		if trimmed != "" {
			bufferedInputLines = append(bufferedInputLines, trimmed)
		}
	}

	// strip the remaining 0A and 0D Linefeed at the end
	for i := 0; i < len(bufferedInputLines); i++ {
		// 0d,0a then again as there have been files observed which had 0a0d (0d0a would be normal)
		bufferedInputLines[i] = strings.Trim(bufferedInputLines[i], string([]byte{0x0A}))
		bufferedInputLines[i] = strings.Trim(bufferedInputLines[i], string([]byte{0x0D}))
		bufferedInputLines[i] = strings.Trim(bufferedInputLines[i], string([]byte{0x0A}))
		bufferedInputLines[i] = strings.Trim(bufferedInputLines[i], string([]byte{0x0D}))
		// fmt.Println(">", bufferedInputLines[i])
	}

	for line := 0; line < len(bufferedInputLines); line++ {
		splitfields := strings.Split(bufferedInputLines[line], separator)

		for i := 0; i < targetStructType.NumField(); i++ {
			currentRecord := targetStructValue.Field(i)
			ftype := targetStructType.Field(i)
			hl7 := ftype.Tag.Get("csv")
			tagsList := removeEmptyStrings(strings.Split(hl7, ","))
			if len(tagsList) >= 1 {

			}
		}
	}

	return nil
}

func EncodeCharsetToUTF8From(charmap *charmap.Charmap, data []byte) ([]byte, error) {
	sr := bytes.NewReader(data)
	e := charmap.NewDecoder().Reader(sr)
	bytes := make([]byte, len(data)*2)
	n, err := e.Read(bytes)
	if err != nil {
		return []byte{}, err
	}
	return bytes[:n], nil
}
