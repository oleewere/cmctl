package cm

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// LoadIniFile load ini configuration file locally
func LoadIniFile(fileName string) (*IniConfiguration, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	cfg := CreateIniConfiguration()
	sections := make(map[string]Section)
	globalSection := cfg.AddSection("global")

	activeSection := globalSection

	fileScanner := bufio.NewScanner(bufio.NewReader(file))
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(strings.TrimSpace(line)) > 0 && !strings.HasPrefix(line, "#") {
			if strings.HasPrefix(line, "[") {
				sectionName := strings.Trim(line, "[]")
				var section Section
				if cfg.SectionExists(sectionName) {
					section = *cfg.GetSection(sectionName)
				} else {
					section = *cfg.AddSection(sectionName)
				}
				activeSection = &section
				sections[sectionName] = section
			} else if strings.Contains(line, "=") {
				section := activeSection
				section.AddSectionKeyValueStr(line)
			} else {
				section := activeSection
				section.AddSectionValue(line)
			}
		}
	}
	if len(globalSection.KeyValueMap) == 0 && len(*globalSection.Values) == 0 {
		delete(sections, "global")
	}
	cfg.Sections = &sections

	return cfg, nil
}

// CreateIniConfiguration create new ini config object
func CreateIniConfiguration() *IniConfiguration {
	sections := make(map[string]Section)
	cfg := IniConfiguration{Sections: &sections}
	return &cfg
}

// WriteIniConfiguration create a new ini file with ini configuration object content - sections are sorted
// but can start with specific sections
func (cfg *IniConfiguration) WriteIniConfiguration(filePath string, startSectionsWith []string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	sections := *cfg.Sections
	startSectionStrs := []string{}
	for _, startSectionName := range startSectionsWith {
		startSection := cfg.GetSection(startSectionName)
		startSectionStrs = append(startSectionStrs, startSection.SectionToString())
		delete(sections, startSectionName)
	}
	sectionStrs := []string{}
	var keys []string
	for key := range sections {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		section := sections[key]
		sectionStrs = append(sectionStrs, section.SectionToString())
	}
	//fmt.Println(startSectionStrs)
	stringToWrite := strings.Join(startSectionStrs, "\n") + "\n" + strings.Join(sectionStrs, "\n")
	_, err = f.WriteString(stringToWrite)
	if err != nil {
		return err
	}
	return nil
}

// SectionToString create a section string (ini like: [section_name]key=value ... with new lines)
func (s *Section) SectionToString() string {
	strResult := ""
	strResult += fmt.Sprintf("[%v]\n", s.Name)
	for _, value := range *s.Values {
		strResult += fmt.Sprintf("%v\n", value)
	}
	for key, value := range s.KeyValueMap {
		strResult += fmt.Sprintf("%v=%v\n", key, value)
	}
	return strResult
}

// AddSection create a new section
func (cfg *IniConfiguration) AddSection(name string) *Section {
	newSlice := make([]string, 0)
	section := Section{Name: name, KeyValueMap: make(map[string]string), Values: &newSlice}
	sections := *cfg.Sections
	sections[name] = section
	return &section
}

// GetSection gather a section by name
func (cfg *IniConfiguration) GetSection(name string) *Section {
	sections := *cfg.Sections
	section := sections[name]
	return &section
}

// AddSectionValue add new section option
func (s *Section) AddSectionValue(value string) *Section {
	values := *s.Values
	values = append(values, value)
	*s.Values = values
	return s
}

// AddSectionKeyValueStr add new section option value (key=value format)
func (s *Section) AddSectionKeyValueStr(keyValue string) *Section {
	splitted := strings.SplitN(keyValue, "=", 2)
	s.AddSectionKeyValue(splitted[0], splitted[1])
	return s
}

// AddSectionKeyValue add new key and value for a section
func (s *Section) AddSectionKeyValue(key string, value string) *Section {
	s.KeyValueMap[key] = value
	return s
}

// SectionExists check that a section exists with a name or not
func (cfg *IniConfiguration) SectionExists(sectionName string) bool {
	sections := *cfg.Sections
	if _, ok := sections[sectionName]; ok {
		return true
	}
	return false
}
