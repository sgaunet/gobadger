# gobadger

Tool to generate badge (svg format).
It has been created to be used for private gitlab repositories. The badge need to be an artifact. See below on how to do.

# Usage

```
$ gobadger -h
Usage of gobadger:
  -c string
        color of badge (default "#5272B4")
  -o string
        output file name (default "badge.svg")
  -t string
        title
  -v string
        Value for the title
```

# In gitlab ci :

Create a stage in the CI to create the badges. They have to be artifacts.

```
build_badges2:
  stage: build
  image: sgaunet/gobadger:latest   # <= replace with the last tag
  script:
      # https://gitlab.com/%{project_path}/-/jobs/artifacts/main/raw/ref.svg?job=build_badges2
      - gobadger -o ref.svg -t godoc -v reference
      # https://gitlab.com/%{project_path}/-/jobs/artifacts/main/raw/badge.svg?job=build_badges2
      - gobadger -o badge.svg -t title -v value -c "#00FF00"
  artifacts:
    name: badge.svg
    paths:
      - badge.svg
      - ref.svg
    expire_in: 2 days
```

And add the badges in the project (settings > General > Badges ):

* Name: Name of the badge
* Link: Link of the button ex: https://gitlab.com/%{project_path}/-/commits/%{default_branch}
* Badge: https://gitlab.com/%{project_path}/-/jobs/artifacts/main/raw/badge.svg?job=build_badges2



