# loop
watch(1) without refreshing the screen

# usage

```bash
loop -n 1s echo hello world
```

```bash
loop -n 1s "date; echo hello world"
```

```bash
loop -n 1s "tput reset; date; echo hello world"
```

