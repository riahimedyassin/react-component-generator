# Components Generation (V0.0.1)

This section describes the components generation flags and behaviour.

## Flags 🚩

The flags should specify a more deterministic and accurate component generation to faciltate and enhance your user experience.
Flags are not mandatory but you should consider them to speed up your developement experience and boost your productivity.
PS : You can combine multiple flags all at once.

### Alias

- -s : With style file
- -t : With test file
- -c : Class component

### Styling 🎨

| Style    | Description                                                                |
| -------- | -------------------------------------------------------------------------- |
| CSS      | **Default** Project styling                                                |
| SCSS     | **rg** will be responsible for **installing** approriate **dependencies**. |
| Tailwind | **rg** will be responsible for **installing** approriate **dependencies**. |

Note that changing your styling configuration after the project initialization could lead to some comptability issues.
Chosing to implement a migration strategy is not decided yet. This could be implmeneted or ignored in future realeses.

### Examples

```bash
rg g c [name] -s -c
```
