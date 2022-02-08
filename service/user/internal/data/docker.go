package data

import (
	"database/sql"
	"fmt"
	"github.com/ory/dockertest/v3"
	log "github.com/sirupsen/logrus"
	"net/url"
	"os"
	"time"
)

func DockerMysql(img, version string) (string, func()) {
	return innerDockerMysql(img, version)
}

func getHostPort(resource *dockertest.Resource, id string) string {
	dockerURL := os.Getenv("DOCKER_HOST")
	if dockerURL == "" {
		return resource.GetHostPort(id)
	}
	u, err := url.Parse(dockerURL)
	if err != nil {
		panic(err)
	}
	return u.Hostname() + ":" + resource.GetPort(id)
}

func innerDockerMysql(img, version string) (string, func()) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 2
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run(img, version, []string{"MYSQL_ROOT_PASSWORD=secret", "MYSQL_ROOT_HOST=%"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	// 5分钟后自动清除
	resource.Expire(600)

	conStr := fmt.Sprintf("root:secret@(localhost:%s)/mysql", resource.GetPort("3306/tcp"))

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		db, err := sql.Open("mysql", conStr)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	return conStr, func() {
		err := resource.Close()
		chk(err)
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
