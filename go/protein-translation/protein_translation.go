package protein

import "errors"

var ErrStop = errors.New("stop")
var ErrInvalidBase = errors.New("invalid base")

var codonMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine",
	"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	var proteins []string
	for i, j := 0, 3; i <= len(rna)-3; i, j = i+3, j+3 {
		var tmp string
		if i == len(rna)-3 {
			tmp = rna[i:]
		} else {
			tmp = rna[i:j]
		}

		pro, err := FromCodon(tmp)
		if err == ErrStop {
			break
		} else if err != nil {
			return nil, err
		}

		proteins = append(proteins, pro)
	}

	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	if protein, ok := codonMap[codon]; !ok {
		return "", ErrInvalidBase
	} else if protein == "STOP" {
		return "", ErrStop
	} else {
		return protein, nil
	}
}
