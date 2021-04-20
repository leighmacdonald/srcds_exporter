# srcds_exporter (w/BattleMetrics)

SRCDS Gameserver Prometheus exporter. This fork includes additional collectors with support for querying
[battlemetrics rankings](https://www.battlemetrics.com/servers/tf2?q=uncletopia&sort=rank).

### Querying BattleMetrics

Note that in order to query (scrape) battlemetrics results we need a full javascript engine. This required 
because the data is loaded in dynamically unfortunately. Because of this, the docker image 
`cdp.Dockerfile` is highly recommended as it is pre-configured to run a headless chrome instance under
a virtual Xorg instance via `Xvfb`. 

If you are not going to use this functionality it is strongly recommended using the [upstream](https://github.com/galexrt/srcds_exporter) 
version of this package instead due to the much larger container sizes and increased resources required for it.

[![Docker Repository on Quay](https://quay.io/repository/galexrt/srcds_exporter/status "Docker Repository on Quay")](https://quay.io/repository/galexrt/srcds_exporter) [![license](https://img.shields.io/github/license/mashape/apistatus.svg)]()

## Compatibility

### Tested Games

* [Garry's Mod](https://store.steampowered.com/app/4000/Garrys_Mod/)
* [Counter-Strike: Source](https://store.steampowered.com/app/240/CounterStrike_Source/)
* [Team Fortress 2](https://store.steampowered.com/app/440)

It may work with newer Source Engine games like [Counter-Strike: Global Offensive](http://store.steampowered.com/app/730/CounterStrike_Global_Offensive/) too, but hasn't been tested too much by the project team.

If you have any issues with a game, please create an issue containing the rcon output of `status` command and we'll look into it.

## Collectors

(*Collectors, the "code" that collects metrics)

Whick collectors are enabled is controlled by the `--collectors.enabled` flag.

### Enabled by default

| Name        | Description                           |
| ----------- | ------------------------------------- |
| playercount | Current player count                  |
| map         | Current map played                    |
| rank        | Battlemetrics global server rankking  |

### Disabled by default

| Name    | Description                                                  |
| ------- | ------------------------------------------------------------ |
| players | Report all players by with their Steam ID label as a metric. |

## Usage

Create the `srcds_exporter` config file (see [srcds.example.yml](srcds.example.yml) for an example). The config file can be named whatever you want, the path to the config must be passed to the `srcds_exporter` through the `-config.file=FILE_PATH` flag (default: `./srcds.yaml` (current directoy file `srcds.yaml`)).

Then just run the `srcds_exporter` binary, through Docker (don't forget to add a mount so the config is available in the container), directly or by having it in your `PATH`.

To get a list of all available flags, use the `--help` flag (`srcds_exporter --help`).

Example output:

```shell
$ srcds_exporter --help
srcds_exporter [FLAGS]
  -collectors.enabled string
    	Comma separated list of active collectors (default "map,playercount,rank")
  -collectors.print
    	If true, print available collectors and exit.
  -config.file string
    	Config file to use. (default "./srcds.yaml")
  -debug
    	Enable debug output
  -help
    	Show help menu
  -version
    	Show version information
  -web.listen-address string
    	The address to listen on for HTTP requests (default ":9137")
  -web.telemetry-path string
    	Path the metrics will be exposed under (default "/metrics")
```

## Docker Image

The Docker image is available from [Docker Hub](https://hub.docker.com/repository/docker/leighmacdonald/srcds_exporter):

* `leighmacdonald/srcds_exporter:latest`
