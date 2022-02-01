# Repoman
Repoman is an utility to manage subrepositories inside [mage](github.com/magefile/mage).
Right now it has four mage commands:

* Pull
  This command clones the repository if doesn't exists and pull the latest changes
  if it does.
* Status
  This command shows the status of the local branches to learn if you need to pull or push
  on any repository.
* Prune
  This command prunes all the remote branches that have already been deleted.
* ChangeAuthor
  Changes the email and name locally for the configured repositories.

In the near future I plan to add a new command to push all the pending commits to the
proper repositories.

## Config file
The projects are defined in a `config.yml` file like the following:

```yaml
author:
  name: Ritho
  email: palvarez@ritho.net

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

## License
Repoman is licensed under the [GNU GPLv3](https://www.gnu.org/licenses/gpl.html).
You should have received a copy of the GNU General Public License along with repoman.
If not, see http://www.gnu.org/licenses/.

<p align="center">
<img src="https://www.gnu.org/graphics/gplv3-127x51.png" alt="GNU GPLv3">
</p>
