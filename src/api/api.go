package api

import (
	"../config"
	consulcli "../consul"
	"../osinfo"
	serfcli "../serf"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

var app *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func attachRoot(app *gin.RouterGroup) {

	/**
	 * Global stats
	 */
	app.GET("/", func(c *gin.Context) {

		c.IndentedJSON(http.StatusOK, gin.H{
			"hostname":       osinfo.Hostname,
			"ipaddr":         osinfo.IPAddress,
			"mem_total":      osinfo.TotalMem,
			"mem_free":       osinfo.FreeMem,
			"mem_user_perc":  osinfo.UsedMem,
			"pid":            os.Getpid(),
			"time":           time.Now(),
			"time_startTime": osinfo.StartTime,
			"time_uptime":    time.Now().Sub(osinfo.StartTime).String()})

	})
}

func Start(cfg config.ApiConfig, cons consulcli.Client, serf serfcli.Client) {

	app = gin.New()
	r := app.Group("/")

	/* attach endpoints */
	attachRoot(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/inventory/:team", func(c *gin.Context) {

		type Host struct {
			Node []string `json:"hosts"`
		}

		type Inventory struct {
			// All     Host `json:"all"`
			Engine  Host `json:"docker_engine"`
			Manager Host `json:"docker_swarm_manager"`
			Worker  Host `json:"docker_swarm_worker"`
		}

		allNodes := []string{}
		allManager := []string{}
		allWorker := []string{}

		status := "alive"
		var tags map[string]string
		tags = make(map[string]string)
		tags["team"] = c.Param("team")

		members, _ := serf.ListMembers(tags, status)
		for _, member := range members {
			allNodes = append(allNodes, member.Name)
			status := osinfo.CheckPort("tcp", member.Name+":2377")
			if status == "Reachable" {
				allManager = append(allManager, member.Name)
			} else {
				allWorker = append(allWorker, member.Name)
			}
		}

		if len(allManager) == 0 {
			allManager = append(allManager, allNodes[0])
			allWorker = append(allWorker[:0], allWorker[1:]...)
		}

		jsons := &Inventory{
			// All: Host{
			// 	Node: allNodes,
			// },
			Engine: Host{
				Node: allNodes,
			},
			Manager: Host{
				Node: allManager,
			},
			Worker: Host{
				Node: allWorker,
			},
		}

		c.JSON(http.StatusOK, jsons)
	})

	r.GET("/member/:team", func(c *gin.Context) {
		status := "alive"
		var tags map[string]string
		tags = make(map[string]string)
		tags["team"] = c.Param("team")

		members, _ := serf.ListMembers(tags, status)
		for _, member := range members {

			var msg struct {
				MemberName    string
				MemberAddress string
			}

			msg.MemberName = member.Name
			msg.MemberAddress = member.Addr.String()

			c.JSON(http.StatusOK, msg)
		}
	})

	r.GET("/members", func(c *gin.Context) {

		members, _ := serf.ListAllMembers()
		for _, member := range members {

			var msg struct {
				MemberName    string
				MemberAddress string
				Team          string
				Role          string
			}

			msg.MemberName = member.Name
			msg.MemberAddress = member.Addr.String()
			msg.Team = member.Tags["team"]
			msg.Role = member.Tags["role"]

			c.JSON(http.StatusOK, msg)
		}
	})

	r.GET("/nodes", func(c *gin.Context) {

		members, _ := cons.ListMembers()
		for _, member := range members {

			var msg struct {
				MemberName    string
				MemberAddress string
				Role          string
			}

			msg.MemberName = member.Name
			msg.MemberAddress = member.Addr
			msg.Role = member.Tags["role"]

			c.JSON(http.StatusOK, msg)
		}
	})

	r.POST("/provision/:env", func(c *gin.Context) {

		// id := c.Query("id")
		// page := c.DefaultQuery("page", "0")
		leader := c.PostForm("leader")
		team := c.PostForm("team")

		switch env := c.Param("env"); env {
		case "aws":
			c.String(http.StatusOK, "Will provision %s environment with leader %s for team %s", env, leader, team)
		case "xen":
			c.String(http.StatusOK, "Will provision %s environment with leader %s for team %s", env, leader, team)

		}

	})

	/* run on port */
	err := app.Run(cfg.Bind)
	if err != nil {
		log.Fatal(err)
	}

}
