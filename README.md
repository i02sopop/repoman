# Repoman
Repoman is an utility to manage subrepositories inside [mage](github.com/magefile/mage).
Right now it has three mage commands:

* Pull
  This command clones the repository if doesn't exists and pull the latest changes
  if it does.
* Status
  This command shows the status of the local branches to learn if you need to pull or push
  on any repository.
* Prune
  This command prunes all the remote branches that have already been deleted.
  
In the near future I plan to add a new command to push all the pending commits to the
proper repositories.

## Config file
The projects are defined in a `config.yml` file like the following:

```yaml
projects:
  repoman:
    repository: "git@github.com:i02sopop/repoman.git"

groups:
  libs:
    mage:
      repository: "git@github.com:magefile/mage.git"
```

The example config file creates the following directory structure:

```
repoman/
  <repoman git repo>
libs/
  mage/
    <mage git repo>
```

The tool doesn't tie you to any git provider (public or private), so it's important
that you put the git url you want to do the clone from.
