# React Component Generator
**React Component Generator** is a CLI to initialize the full folder structure for your React Project Component.


* ## Install 
1. For Developpment Purpuses.
###### This CLI Tool should be better installed globaly so you could simply access the command line quickly.
```
npm install react-component-generator --save-dev -g
```
2. General Purpuses
###### This CLI Tool should be better installed globaly so you could simply access the command line quickly.
```
npm install react-component-generator -g
```
* ## Config File
  The **Config File** under the name of *generator.config.json* is an important file to setup your custom file generator.
  * ### Intialize the config file :
    To Initialize the config file , write down the following command :
    ```
    react-config 
    ```
  * You can specify :
  1. React Template
     * Javascript Template
     * Typescript Template
  2. Style Template
     * CSS
     * SCSS
     * SASS
  3. Destination
     * All of the destination files should be under the **src** folder. You have to type only the folder name. If it **doesn't** exist it will be generated.
  4. Component Type :
     * Functionnal Component
     * Class Component
### Example 
This config file will generate for you a file under the **components** folder. The template will be **JavaScript** *(.jsx file)*. The styling will be **SCSS** file and the type of component will be **Functionnal** 
```json
    {
        "template" : "javascript",
        "style" : "scss"
        "dist" : "components",
        "c" : "f"
    }
```
---

### Properties and Values
* Template :
  
| Property | Values    | Output    |
| :-----: | :---: | :---: |
| Template | javascript   | .jsx file   |
|  | typescript   | .tsx file   |

* Dist :
  
| Property | Values    | Output    |
| :-----: | :---: | :---: |
| Dist | Example   | Generate Component Under : /src/Example   |

* Style :
  
| Property | Values    | Output    |
| :-----: | :---: | :---: |
| Style | css   | .css file   |
|  | scss   | .scss file   |
|  | sass   | .sass file   |

* Component :
  
| Property | Values    | Output    |
| :-----: | :---: | :---: |
| c | c   | Class Component  |
| c | f   | Functionnal Component  |

**NOTICE** : The generator file comes with a **JavaScript CSS Functionnal Component** Default Values.

## Dynamic Generation 
You can specify the **Name of Component** and generated file type directly from your command line 

1. Component Name :
   To Specify the component name , you have to specify the **--name** field value.
   The **Component Name** will be automatically parsed into a valid Component Name :
     * Starts With a Capital Letter
     * Should Not hold spaces (spaces will be replaced with dashes)
     * Should not hold any extension
   ---
   ### Example 1:
   ```
   react-comp -name="ExampleComponent"
   ```
   ---
    ### Example 2:
   ```
   react-comp -name="example component"
   ```
   **Result** : Example-component
   
   ---
2. Component Proprities :
   To set up a custom component dynamically seperatly from the **generator.config.json** setting , you could do that simply by setting up the field name and the value for it.
   ```
   react-comp --name="" --template="" --c="" --style="" --dist=""
   ```
    ## Properties :
     
    | Property | Default    | Required    | Descreption    |
    | :-----: | :---: | :---: | :---: |
    | name | - | true  | Component Name |
    | style | css | false  | Style Type  |
    | dist | components | false  | Component Distination  |
    | c | f | false  | Component Type  |

---
## V 1.0.0 
This version is still under developpment . It may hold multiple bugs , I'll be working on imporving it . 
Developped By : 
[**Riahi Yassin**](https://www.linkedin.com/in/riahi-mohamed-yassin/ "LinkedIn Profile")
---
   
   






