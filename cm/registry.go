package cm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

const (
	cmServerJSONFileName           = "cm_servers.json"
	connectionProfilesJSONFileName = "connection_profiles.json"
)

// CreateCMRegistryDb initialize cmctl database
func CreateCMRegistryDb() {
	cmServerJSONFile := getJSONDbFile(cmServerJSONFileName)
	connectionProfileJSONFile := getJSONDbFile(connectionProfilesJSONFileName)
	if !Exists(cmServerJSONFile) {
		cmServerRegistries := make([]CMServer, 0)
		cmServerJSON, _ := json.Marshal(cmServerRegistries)
		err := ioutil.WriteFile(cmServerJSONFile, cmServerJSON, 0644)
		checkErr(err)
	}
	if !Exists(connectionProfileJSONFile) {
		connectionProfiles := make([]ConnectionProfile, 0)
		connectionProfilesJSON, _ := json.Marshal(connectionProfiles)
		err := ioutil.WriteFile(connectionProfileJSONFile, connectionProfilesJSON, 0644)
		checkErr(err)
	}
}

// DropCMRegistryRecords drop all CM server entries from cmctl database
func DropCMRegistryRecords() {
	cmServerRegistries := make([]CMServer, 0)
	WriteCMServerEntries(cmServerRegistries)
}

// DropConnectionProfileRecords drop all connection profile from cmctl database
func DropConnectionProfileRecords() {
	connectionProfiles := make([]ConnectionProfile, 0)
	WriteConnectionProfileEntries(connectionProfiles)
}

// ListCMRegistryEntries get all CM server registries from cmctl database
func ListCMRegistryEntries() []CMServer {
	cmServerJSONFile := getJSONDbFile(cmServerJSONFileName)
	file, err := ioutil.ReadFile(cmServerJSONFile)
	checkErr(err)
	cmRegistries := make([]CMServer, 0)
	json.Unmarshal(file, &cmRegistries)
	return cmRegistries
}

// ListConnectionProfileEntries get all CM server registries from cmctl database
func ListConnectionProfileEntries() []ConnectionProfile {
	connectionProfileJSONFile := getJSONDbFile(connectionProfilesJSONFileName)
	file, err := ioutil.ReadFile(connectionProfileJSONFile)
	checkErr(err)
	connectionProfiles := make([]ConnectionProfile, 0)
	json.Unmarshal(file, &connectionProfiles)
	return connectionProfiles
}

// GetCMEntryID get CM server entry id if the id exists
func GetCMEntryID(id string) string {
	cmEntries := ListCMRegistryEntries()
	cmEntryID := ""
	if len(cmEntries) > 0 {
		for _, cmEntry := range cmEntries {
			if cmEntry.Name == id {
				cmEntryID = cmEntry.Name
			}
		}
	}
	return cmEntryID
}

// GetConnectionProfileEntryID get connection profile entry id if the id exists
func GetConnectionProfileEntryID(id string) string {
	connectionProfiles := ListConnectionProfileEntries()
	connectionProfileID := ""
	if len(connectionProfiles) > 0 {
		for _, connectionProfileEntry := range connectionProfiles {
			if connectionProfileEntry.Name == id {
				connectionProfileID = connectionProfileEntry.Name
			}
		}
	}
	return connectionProfileID
}

// RegisterNewCMEntry create new CM server registry entry in cmctl database
func RegisterNewCMEntry(id string, hostname string, port int, protocol string, username string, password string, useGateway bool, apiVersion int) {
	checkID := GetCMEntryID(id)
	if len(checkID) > 0 {
		alreadyExistMsg := fmt.Sprintf("CM Server with id '%s' is already defined in CM server registry DB.", checkID)
		fmt.Println(alreadyExistMsg)
		os.Exit(1)
	}
	cmServerEntries := ListCMRegistryEntries()
	newCmServerEntry := CMServer{Name: id, Hostname: hostname, Port: port, Protocol: protocol, Username: username,
		Password: password, Active: true, UseGateway: useGateway, APIVersion: apiVersion}
	cmServerEntries = append(cmServerEntries, newCmServerEntry)
	WriteCMServerEntries(cmServerEntries)
}

// UpdateCMEntry update CM server registry entry in cmctl database
func UpdateCMEntry(id string, hostname string, port int, protocol string, username string, password string, useGateway bool, apiVersion int, connectionProfile string) {
	checkID := GetCMEntryID(id)
	if len(checkID) == 0 {
		notExistMsg := fmt.Sprintf("CM Server with id '%s' does not exist in CM server registry DB.", checkID)
		fmt.Println(notExistMsg)
		os.Exit(1)
	}
	updatedCmServerEntry := CMServer{Name: id, Hostname: hostname, Port: port, Protocol: protocol, Username: username, Password: password,
		Active: true, UseGateway: useGateway, APIVersion: apiVersion, ConnectionProfile: connectionProfile}

	cmServers := ListCMRegistryEntries()
	newCmServers := make([]CMServer, 0)
	if len(cmServers) > 0 {
		for index := range cmServers {
			if cmServers[index].Name != id {
				newCmServers = append(newCmServers, cmServers[index])
			} else {
				newCmServers = append(newCmServers, updatedCmServerEntry)
			}
		}
	}
	WriteCMServerEntries(newCmServers)
}

// RegisterNewConnectionProfile create new connection profile entry in cmctl database
func RegisterNewConnectionProfile(id string, keyPath string, port int, username string) {
	checkID := GetConnectionProfileEntryID(id)
	if len(checkID) > 0 {
		alreadyExistMsg := fmt.Sprintf("Connection profile with id '%s' is already defined as a profile entry", checkID)
		fmt.Println(alreadyExistMsg)
		os.Exit(1)
	}
	connectionProfiles := ListConnectionProfileEntries()
	newConnectionProfile := ConnectionProfile{Name: id, KeyPath: keyPath, Port: port, Username: username}
	connectionProfiles = append(connectionProfiles, newConnectionProfile)
	WriteConnectionProfileEntries(connectionProfiles)
}

// DeRegisterCMEntry remove an CM server enrty by id
func DeRegisterCMEntry(id string) {
	cmServers := ListCMRegistryEntries()
	newCmServers := make([]CMServer, 0)
	if len(cmServers) > 0 {
		for index := range cmServers {
			if cmServers[index].Name != id {
				newCmServers = append(newCmServers, cmServers[index])
			}
		}
	}
	WriteCMServerEntries(newCmServers)
}

// DeRegisterConnectionProfile remove a connection profile by id
func DeRegisterConnectionProfile(id string) {
	connectionProfiles := ListConnectionProfileEntries()
	newConnectionProfiles := make([]ConnectionProfile, 0)
	if len(connectionProfiles) > 0 {
		for index := range connectionProfiles {
			if connectionProfiles[index].Name != id {
				newConnectionProfiles = append(newConnectionProfiles, connectionProfiles[index])
			}
		}
	}
	WriteConnectionProfileEntries(newConnectionProfiles)
}

// GetActiveCM get the active CM server registry from cmctl database (should be only one)
func GetActiveCM() CMServer {
	cmServers := ListCMRegistryEntries()
	var result CMServer
	if len(cmServers) > 0 {
		for _, cmServerEntry := range cmServers {
			if cmServerEntry.Active {
				result = cmServerEntry
			}
		}
	}
	return result
}

// GetCMByID get the CM server registry from cmctl database by id
func GetCMByID(searchID string) CMServer {
	cmServers := ListCMRegistryEntries()
	var result CMServer
	if len(cmServers) > 0 {
		for _, cmServerEntry := range cmServers {
			if cmServerEntry.Name == searchID {
				result = cmServerEntry
			}
		}
	}
	return result
}

// GetConnectionProfileByID get the connection profile from cmctl database by id
func GetConnectionProfileByID(searchID string) ConnectionProfile {
	connectionProfiles := ListConnectionProfileEntries()
	var result ConnectionProfile
	if len(connectionProfiles) > 0 {
		for _, connectionProfileEntry := range connectionProfiles {
			if connectionProfileEntry.Name == searchID {
				result = connectionProfileEntry
			}
		}
	}
	return result
}

// SetProfileIDForCMEntry attach a connection profile to a specific CM server entry
func SetProfileIDForCMEntry(cmEntryID string, profileId string) {
	cmServers := ListCMRegistryEntries()
	if len(cmServers) > 0 {
		for index := range cmServers {
			if cmServers[index].Name == cmEntryID {
				cmServers[index].ConnectionProfile = profileId
			}
		}
	}
	WriteCMServerEntries(cmServers)
}

// ActiveCMRegistry turn on active status on selected CM registry
func ActiveCMRegistry(id string) {
	checkID := GetCMEntryID(id)
	if len(checkID) == 0 {
		alreadyExistMsg := fmt.Sprintf("Not found CM server registry  with id '%s'.", checkID)
		fmt.Println(alreadyExistMsg)
		os.Exit(1)
	}
	cmServers := ListCMRegistryEntries()
	if len(cmServers) > 0 {
		for index := range cmServers {
			if cmServers[index].Name == id {
				cmServers[index].Active = true
			} else {
				cmServers[index].Active = false
			}
		}
	}
	WriteCMServerEntries(cmServers)
}

// DeactiveAllCMRegistry turn off active status on all CM registries
func DeactiveAllCMRegistry() {
	cmServers := ListCMRegistryEntries()
	if len(cmServers) > 0 {
		for index := range cmServers {
			cmServers[index].Active = false
		}
	}
	WriteCMServerEntries(cmServers)
}

// WriteCMServerEntries write CM server entries to the CM server registry json file
func WriteCMServerEntries(cmServers []CMServer) {
	cmServerJSON, _ := json.Marshal(cmServers)
	cmServerJSONFile := getJSONDbFile(cmServerJSONFileName)
	err := ioutil.WriteFile(cmServerJSONFile, FormatJSON(cmServerJSON).Bytes(), 0600)
	checkErr(err)
}

// WriteConnectionProfileEntries write connection profile entries to the connection profile registry json file
func WriteConnectionProfileEntries(connectionProfiles []ConnectionProfile) {
	connectionProfilesJSON, _ := json.Marshal(connectionProfiles)
	connectionProfilesJSONFile := getJSONDbFile(connectionProfilesJSONFileName)
	err := ioutil.WriteFile(connectionProfilesJSONFile, FormatJSON(connectionProfilesJSON).Bytes(), 0600)
	checkErr(err)
}

// FormatJSON format json file
func FormatJSON(b []byte) *bytes.Buffer {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &out
}

func getJSONDbFile(file string) string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	home := usr.HomeDir
	cmManagerFolder := path.Join(home, ".cmctl")
	if _, err := os.Stat(cmManagerFolder); os.IsNotExist(err) {
		os.Mkdir(cmManagerFolder, os.ModePerm)
	}
	return path.Join(cmManagerFolder, file)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
