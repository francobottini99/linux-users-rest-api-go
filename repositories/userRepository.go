package repository

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"time"

	model "github.com/ICOMP-UNC/2023---soii---laboratorio-6-FrancoNB/models"
	"golang.org/x/crypto/ssh"
)

func UserGetUID(username string) (int, error) {
	u, err := user.Lookup(username)

	if err != nil {
		return 0, err
	}

	uid, err := strconv.Atoi(u.Uid)

	if err != nil {
		return 0, err
	}

	return uid, nil
}

func UserGetRegistrationTime(username string) (time.Time, error) {
	_, err := exec.Command("id", "-u", username).Output()

	if err != nil {
		return time.Time{}, err
	}

	statCmd := exec.Command("stat", "-c", "%Y", "/home/"+username)
	output, err := statCmd.Output()

	if err != nil {
		return time.Time{}, err
	}

	creationTime, err := strconv.ParseInt(strings.TrimSpace(string(output)), 10, 64)

	if err != nil {
		return time.Time{}, err
	}

	registrationTime := time.Unix(creationTime, 0)

	return registrationTime, nil
}

func UserCreate(user model.User) (model.User, error) {
	var result model.User

	cmd := exec.Command("useradd", "-m", "-s", "/bin/bash", user.Username)
	err := cmd.Run()

	if err != nil {
		log.Println(err)
		return result, fmt.Errorf("Could not create user '%s': %v", user.Username, err)
	}

	cmd = exec.Command("bash", "-c", fmt.Sprintf("echo '%s:%s' | chpasswd", user.Username, user.Password))
	err = cmd.Run()

	if err != nil {
		log.Println(err)
		return result, fmt.Errorf("Could not set user password '%s': %v", user.Username, err)
	}

	uid, err := UserGetUID(user.Username)

	if err != nil {
		log.Println(err)
		return result, fmt.Errorf("Could not get user uid '%s': %v", user.Username, err)
	}

	err = os.Chown("/home/"+user.Username, uid, -1)

	if err != nil {
		log.Println(err)
		return result, fmt.Errorf("Could not change owner of user's home directory '%s': %v", user.Username, err)
	}

	registrationTime, err := UserGetRegistrationTime(user.Username)

	if err != nil {
		log.Println(err)
		return result, fmt.Errorf("Could not get user registration time '%s': %v", user.Username, err)
	}

	result.Id = uid
	result.Username = user.Username
	result.Password = ""
	result.CreateAt = registrationTime.Format("2006-01-02 15:04:05")

	return result, nil
}

func UserDelete(username string) error {
	cmd := exec.Command("userdel", "-r", username)
	err := cmd.Run()

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Could not delete user '%s': %v", username, err)
	}

	return nil
}

func UserGet(username string) (model.User, error) {
	var result model.User

	u, err := user.Lookup(username)

	if err != nil {
		log.Println(err)
		return result, fmt.Errorf("User '%s' does not exist", username)
	}

	uid, err := strconv.Atoi(u.Uid)

	if err != nil {
		log.Println(err)
		return result, fmt.Errorf("Could not get user UID '%s': %v", username, err)
	}

	registrationTime, err := UserGetRegistrationTime(username)

	if err != nil {
		log.Println(err)
		return result, fmt.Errorf("Could not get user registration time '%s': %v", username, err)
	}

	result.Id = uid
	result.Username = u.Username
	result.CreateAt = registrationTime.Format("2006-01-02 15:04:05")

	return result, nil
}

func UserGetAll() ([]model.User, error) {
	cmd := exec.Command("sh", "-c", "grep '^[^:]*:[^:]*:[0-9]\\{4,\\}:' /etc/passwd | awk -F: '$3 >= 1000 && $1 != \"nobody\" {print $1}'")
	output, err := cmd.Output()

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Could not get user list: %v", err)
	}

	var userList []model.User

	users := strings.Split(string(output), "\n")

	for _, username := range users {
		if username != "" {
			var user model.User

			uid, err := UserGetUID(username)

			if err != nil {
				log.Println(err)
				return nil, fmt.Errorf("Could not get user UID '%s': %v", username, err)
			}

			registrationTime, err := UserGetRegistrationTime(username)

			if err != nil {
				log.Println(err)
				return nil, fmt.Errorf("Could not get user registration time '%s': %v", username, err)
			}

			user.Id = uid
			user.Username = username
			user.Password = ""
			user.CreateAt = registrationTime.Format("2006-01-02 15:04:05")

			userList = append(userList, user)
		}
	}

	return userList, nil
}

func UserValidateCredentials(username string, password string) error {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", "localhost:22", config)

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Invalid password for user '%s'", username)
	}

	defer client.Close()

	return nil
}
