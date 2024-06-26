package simpletaxonomy

import (
	"encoding/json"
	"fmt"
	"os"
)

type SimplifiedTaxonomy struct {
    Taxonomies []TaxL1 `json:"simplified-taxonomy"`
}

type TaxL1 struct {
	Name       string  `json:"li-taxonomy"`
	Slug       string  `json:"slug"`
	Definition string  `json:"definition"`
	Children   []TaxL2 `json:"children"`
}

type TaxL2 struct {
	Name       string `json:"l2-taxonomy"`
	Slug       string `json:"slug"`
	Definition string `json:"definition"`
	Squad      []string `json:"squad"`
}

func Load() error {
    content, err := os.ReadFile("./simplified-taxonomy.json")
    if err != nil {
        return err
    }

    var taxonomies SimplifiedTaxonomy
    err = json.Unmarshal(content, &taxonomies)
    if err != nil {
        return err
    }

    fmt.Println(taxonomies)

    return nil
}

