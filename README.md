# React Component Generator
**React Component Generator** is a CLI to generate Components for your React Project.


* ## Install 
###### This CLI Tool should be better installed globally, so you could simply access the command line quickly.
```
npm install rcg-cli-tool -g
```
* ## Config File
  The **Config File** under the name of *generator.config.json* is an important file to set up your custom file generator.
  * ### Initialize the config file :
    To Initialize the config file , write down the following command :
    ```
    react-config 
    ```
  * You can specify :
  1. React Template : 
     * Javascript Template
     * Typescript Template
  2. Style Template
     * CSS
     * SCSS
     * SASS
  3. Destination
     * All the destination files should be under the **src** folder. You have to type only the folder name. If it **doesn't** exist it will be generated.
  4. Component Type :
     * Functional Component
     * Class Component
### Example 
This config file will generate for you a file under the **components** folder. The template will be **JavaScript** *(.jsx file)*. The styling will be **SCSS** file and the type of component will be **Functional** 
```json
    {
        "template" : "-js",
        "style" : "-scss",
        "dist" : "components",
        "type" : "function"
    }
```
---

### Properties and Values
* Template :
  
| Property |   Values   |  Output   |
|:--------:|:----------:|:---------:|
| Template | -js | .jsx file |
|          | -ts | .tsx file |

* Dist :
  
| Property | Values  |                 Output                  |
|:--------:|:-------:|:---------------------------------------:|
|   Dist   | Example | Generate Component Under : /src/Example |

* Style :
  
| Property | Values |   Output   |
|:--------:|:------:|:----------:|
|  Style   |  -css   | .css file  |
|          |  -scss  | .scss file |
|          |  -sass  | .sass file |
|          |  -less  | .less file |

* Component :
  
| Property |  Values  |        Output        |
|:--------:|:--------:|:--------------------:|
|   type   |  class   |   Class Component    |
|          | function | Functional Component |

**NOTICE** : The generator file initialy have **JavaScript CSS Functional Component** Default Values.

## Dynamic Generation 
You can specify the **Name of Component** and generated file type directly from your command line 

1. Component Name :
   To Specify the component name , you have to type it directelly after the **react** command.
   The **Component Name** will be **automatically parsed** into a valid Component Name :
     * Starts With a Capital Letter
     * Should Not hold spaces (spaces will be replaced with dashes)
     * Should not hold any extension
   ---
   ### Example 1:
   ```
   react ExampleComponent
   ```
   **Result** : ExampleComponent
   ---
    ### Example 2:
   ```
   react example-component
   ```
   **Result** : Example-component
   
   ---
2. Component Proprieties :
   To set up a custom component dynamically separately from the **generator.config.json** setting , you could do that simply by setting up the field name and the value for it.
   ```
   react FolderName -template (-ts || -js) -style (-css || -scss || -sass || -less)
   ```
    ## Properties :
     
    | Property |  Default   | Required |      Description      |
    |:--------:|:----------:|:--------:|:---------------------:|
    |   name   |     -      |   true   |    Component Name     |
    |  style   |    -css     |  false   |      Style Type       |
    |   dist   | components |  false   | Component Destination |
    |   type   |  function  |  false   |    Component Type     |

---
## V 2.0.0 Features 
1. CLI TOOL Name 
   For simplicity , the command line tool had been changed from *react-comp* to **react**

2. Test Files 
   From now on , you can get rid of the Test files without having to delete them manually everytime.
   You can specify the **test** field in the **generator.config.json** file to false you'll not get any testing files anymore.

   ```json
      {
         "...options":"...values",
         "test":false
      }

   ```

---
## V 2.1.0 
This version is still under development . It may hold multiple bugs , I'll be working on improving it . 

I had recently changed to CLI global command to **rg** instead of react.
I had also added Less as a style sheet option.

### Previous Version : 
```
react FolderName -ts -scss
```
This Approach wasn't really fast enough and it any typing error may cause you an unexcpeted error behaviour 

### Latest Version : 
```
rg FolderName -ts -scss
```
This approach is very efficient in terms of time saving.
It is more of inspiration from the **ng** command line tool from Angular .
Yet , issues to be solved but You could use it Safelly in your projects.

##### *Note*  
I tried to write a cleaner code compared to the previous version so I could better maintain it.

---

# Developed By : **Riahi Yassin**
[**LinkedIn**](https://www.linkedin.com/in/riahi-mohamed-yassin/ "LinkedIn Profile")
[**Instagram**](https://www.instagram.com/riahi__yassin/ "Instagram Profile")

---

   
   






