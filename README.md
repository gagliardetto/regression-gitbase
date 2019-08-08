# regression-gitbase

**regression-gitbase** is a tool than runs different versions of `gitbase` and compares its resource consumption.

```
Usage:
  regression [OPTIONS]

gitbase regression tester.

This tool executes several versions of gitbase and compares query times and
resource usage. There should be at least two versions specified as arguments in
the following way:

* v0.12.1 - release name from github (https://github.com/src-d/gitbase/releases).
The binary will be downloaded.
* latest - latest release from github. The binary will be downloaded.
* remote:master - any tag or branch from gitbase repository. The binary will be
built automatically.
* local:fix/some-bug - tag or branch from the repository in the current directory.
The binary will be built.
* local:HEAD - current state of the repository. Binary is built.
* pull:266 - code from pull request #266 from gitbase repo. Binary is built.
* /path/to/gitbase - a gitbase binary built locally.

The repositories and downloaded/built gitbase binaries are cached by default in
"repos" and "binaries" repositories from the current directory.


Application Options:
      --binaries=     Directory to store binaries (default: binaries) [$REG_BINARIES]
      --repos=        Directory to store repositories (default: repos) [$REG_REPOS]
      --url=          URL to the tool repo [$REG_GITURL]
      --gitport=      Port for local git server (default: 9418) [$REG_GITPORT]
      --repos-file=   YAML file with the list of repos [$REG_REPOS_FILE]
  -c, --complexity=   Complexity of the repositories to test (default: 1) [$REG_COMPLEXITY]
  -n, --repeat=       Number of times a test is run (default: 3) [$REG_REPEAT]
      --show-repos    List available repositories to test
  -t, --token=        Token used to connect to the API [$REG_TOKEN]
      --csv           save csv files with last result
      --prom          store latest results to prometheus
      --prom-address= prometheus pushgateway address [$PROM_ADDRESS]
      --prom-job=     prometheus job [$PROM_JOB]
      --ci-branch=    branch env [$GIT_BRANCH]
      --ci-commit=    commit env [$GIT_COMMIT]

Help Options:
  -h, --help        Show this help message
```

## License

Licensed under the terms of the Apache License Version 2.0. See the `LICENSE`
file for the full license text.

