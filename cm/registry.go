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
	cmServerJsonFileName           = "cm_servers.json"
	connectionProfilesJsonFileName = "connection_profiles.json"
)

// CreateCMRegistryDb initialize cmctl database
func CreateCMRegistryDb() {
	cmServerJsonFile := getJsonDbFile(cmServerJsonFileName)
	connectionProfileJsonFile := getJsonDbFile(connectionProfilesJsonFileName)
	if !exists(cmServerJsonFile) {
		cmServerRegistries := make([]CMServer, 0)
		cmServerJson, _ := json.Marshal(cmServerRegistries)
		err := ioutil.WriteFile(cmServerJsonFile, cmServerJson, 0644)
		checkErr(err)
	}
	if !exists(connectionProfileJsonFile) {
		connectionProfiles := make([]ConnectionProfile, 0)
		connectionProfilesJson, _ := json.Marshal(connectionProfiles)
		err := ioutil.WriteFile(connectionProfileJsonFile, connectionProfilesJson, 0644)
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
	cmServerJsonFile := getJsonDbFile(cmServerJsonFileName)
	file, err := ioutil.ReadFile(cmServerJsonFile)
	checkErr(err)
	cmRegistries := make([]CMServer, 0)
	json.Unmarshal(file, &cmRegistries)
	return cmRegistries
}

// ListConnectionProfileEntries get all CM server registries from cmctl database
func ListConnectionProfileEntries() []ConnectionProfile {
	connectionProfileJsonFile := getJsonDbFile(connectionProfilesJsonFileName)
	file, err := ioutil.ReadFile(connectionProfileJsonFile)
	checkErr(err)
	connectionProfiles := make([]ConnectionProfile, 0)
	json.Unmarshal(file, &connectionProfiles)
	return connectionProfiles
}

// GetCMEntryId get CM server entry id if the id exists
func GetCMEntryId(id string) string {
	cmEntries := ListCMRegistryEntries()
	cmEntryId := ""
	if len(cmEntries) > 0 {
		for _, cmEntry := range cmEntries {
			if cmEntry.Name == id {
				cmEntryId = cmEntry.Name
			}
		}
	}
	return cmEntryId
}

// GetConnectionProfileEntryId get connection profile entry id if the id exists
func GetConnectionProfileEntryId(id string) string {
	connectionProfiles := ListConnectionProfileEntries()
	connectionProfileId := ""
	if len(connectionProfiles) > 0 {
		for _, connectionProfileEntry := range connectionProfiles {
			if connectionProfileEntry.Name == id {
				connectionProfileId = connectionProfileEntry.Name
			}
		}
	}
	return connectionProfileId
}

// RegisterNewCMEntry create new CM server registry entry in cmctl database
func RegisterNewCMEntry(id string, hostname string, port int, protocol string, username string, password string, useGateway bool, apiVersion int) {
	checkId := GetCMEntryId(id)
	if len(checkId) > 0 {
		alreadyExistMsg := fmt.Sprintf("Registry with id '%s' is already defined as a registry entry", checkId)
		fmt.Println(alreadyExistMsg)
		os.Exit(1)
	}
	cmServerEntries := ListCMRegistryEntries()
	newCmServerEntry := CMServer{Name: id, Hostname: hostname, Port: port, Protocol: protocol, Username: username,
		Password: password, Active: true, UseGateway: useGateway, ApiVersion: apiVersion}
	cmServerEntries = append(cmServerEntries, newCmServerEntry)
	WriteCMServerEntries(cmServerEntries)
}

// UpdateCMEntry update CM server registry entry in cmctl database
func UpdateCMEntry(id string, hostname string, port int, protocol string, username string, password string, useGateway bool, apiVersion int, connectionProfile string) {
	checkId := GetCMEntryId(id)
	if len(checkId) == 0 {
		notExistMsg := fmt.Sprintf("Registry with id '%s' does not exist in CM server registry DB.", checkId)
		fmt.Println(notExistMsg)
		os.Exit(1)
	}
	updatedCmServerEntry := CMServer{Name: id, Hostname: hostname, Port: port, Protocol: protocol, Username: username, Password: password,
		Active: true, UseGateway: useGateway, ApiVersion: apiVersion, ConnectionProfile: connectionProfile}

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
	checkId := GetConnectionProfileEntryId(id)
	if len(checkId) > 0 {
		alreadyExistMsg := fmt.Sprintf("Connection profile with id '%s' is already defined as a profile entry", checkId)
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

// GetCMById get the CM server registry from cmctl database by id
func GetCMById(searchId string) CMServer {
	cmServers := ListCMRegistryEntries()
	var result CMServer
	if len(cmServers) > 0 {
		for _, cmServerEntry := range cmServers {
			if cmServerEntry.Name == searchId {
				result = cmServerEntry
			}
		}
	}
	return result
}

// GetConnectionProfileById get the connection profile from cmctl database by id
func GetConnectionProfileById(searchId string) ConnectionProfile {
	connectionProfiles := ListConnectionProfileEntries()
	var result ConnectionProfile
	if len(connectionProfiles) > 0 {
		for _, connectionProfileEntry := range connectionProfiles {
			if connectionProfileEntry.Name == searchId {
				result = connectionProfileEntry
			}
		}
	}
	return result
}

// SetProfileIdForCMEntry attach a connection profile to a specific CM server entry
func SetProfileIdForCMEntry(cmEntryId string, profileId string) {
	cmServers := ListCMRegistryEntries()
	if len(cmServers) > 0 {
		for index := range cmServers {
			if cmServers[index].Name == cmEntryId {
				cmServers[index].ConnectionProfile = profileId
			}
		}
	}
	WriteCMServerEntries(cmServers)
}

// ActiveCMRegistry turn on active status on selected CM registry
func ActiveCMRegistry(id string) {
	checkId := GetCMEntryId(id)
	if len(checkId) == 0 {
		alreadyExistMsg := fmt.Sprintf("Not found CM server registry  with id '%s'.", checkId)
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
	cmServerJson, _ := json.Marshal(cmServers)
	cmServerJsonFile := getJsonDbFile(cmServerJsonFileName)
	err := ioutil.WriteFile(cmServerJsonFile, FormatJson(cmServerJson).Bytes(), 0600)
	checkErr(err)
}

// WriteConnectionProfileEntries write connection profile entries to the connection profile registry json file
func WriteConnectionProfileEntries(connectionProfiles []ConnectionProfile) {
	connectionProfilesJson, _ := json.Marshal(connectionProfiles)
	connectionProfilesJsonFile := getJsonDbFile(connectionProfilesJsonFileName)
	err := ioutil.WriteFile(connectionProfilesJsonFile, FormatJson(connectionProfilesJson).Bytes(), 0600)
	checkErr(err)
}

// FormatJson format json file
func FormatJson(b []byte) *bytes.Buffer {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &out
}

func getJsonDbFile(file string) string {
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

// Exists reports whether the named file or directory exists.
func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
