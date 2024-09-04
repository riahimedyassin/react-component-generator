# React Component Generator

**React Component Generator** is a CLI to generate Components for your React Project.

- ## Install

###### This CLI Tool must be installed globally.

```
npm install rcg-cli-tool@latest -g
```

- ## Config File
  The **Config File** is named _rg.config.json_ .
  ### Initialize the config file :
  To Initialize the config file , write down the following command :
  ```
  rg --init
  ```

### Example

This config file will generate for you a file under the **components** folder. The template will be **JavaScript** _(.jsx file)_. The styling will be **SCSS** file and the type of component will be **Functional**

```json
{
  "template": "jsx",
  "style": "scss",
  "dist": "components",
  "type": "rfce"
}
```

---

### Properties and Values

- Template :

| Property | Values |  Output   |
| :------: | :----: | :-------: |
| Template |  jsx   | .jsx file |
|          |  tsx   | .tsx file |

- Dist :

| Property | Values  |                 Output                 |
| :------: | :-----: | :------------------------------------: |
|   Dist   | Example | Generate Component Under : **Example** |

- Style :

| Property | Values |   Output   |
| :------: | :----: | :--------: |
|  Style   |  css   | .css file  |
|          |  scss  | .scss file |
|          |  sass  | .sass file |
|          |  less  | .less file |

- Component :

| Property | Values |        Output        |
| :------: | :----: | :------------------: |
|   type   |  rce   |   Class Component    |
|          |  rfce  | Functional Component |

**_Default_** : The config file initialy configured to : **JavaScript CSS Functional Component** .

## Dynamic Generation

You can specify the **Name of Component** and generated file type directly from your command line

1. **Component Generation :**
   You can specify custom properties for your components dynamically.
   ```
   rg component_name (ts) (sass) (test)
   ```

---

