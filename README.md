# multi-lang-repo

A demo repository with simple hello-world apps written in four languages, each with its
own package manager and a few real third-party dependencies (useful as SCA scan targets).

| App | Language | Package Manager | Manifest |
|-----|----------|------------------|----------|
| [python-app](python-app/) | Python | pip | `requirements.txt` |
| [java-app](java-app/) | Java | Maven | `pom.xml` |
| [javascript-app](javascript-app/) | JavaScript | npm | `package.json` |
| [go-app](go-app/) | Go | Go modules | `go.mod` |

Each app's README has run instructions. Dependency versions are intentionally pinned
(some to older releases) to give an SCA scanner something to find.
