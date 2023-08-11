#!/usr/bin/env node
'use strict';
/**
 * @requires fs
 * @requires path
 * @requires yargs
 * @requires "./utils/UppedCase"
 * @requires "./utils/checkTemplate"
 * @requires "./generator.config.json"
 * @requires "./utils/constants"
 * @description CLI TOOL that initialize the React Project Folder Structure {Folder->{JSX/TSX File,CSS/SCSS/SASS File,Testing File}}
 *
 *
 * */


const fs = require("fs")
const path = require("path")
const yargs = require("yargs")
const { argv } = yargs(process.argv)
const UpperCase = require("./utils/UppedCase")
const createFile = require("./utils/createFile");
const checkTemplate = require("./utils/checkTemplate");
const generator = JSON.parse(fs.readFileSync("./generator.config.json", { encoding: "utf-8" }).toString())
const { scriptPossible, stylePossible, componentPossible } = require("./utils/constants")
let hasError = false
const setHasError = (bool) => {
    hasError = bool
}



//Initial Values from the generator
if (!generator) return console.log("Please head to the initiale folder to where the generator.config.json file is located")
const { style, template, type, dist } = generator





if (!argv.name) return console.log("All fields are mandatory")

//Where the folder will be created
const destinationFile = argv.folder || dist


//The created Folder path and name
const folderPath = `./src/${destinationFile}`;
const folderName = argv.name.trim().split(" ").join("-")
//Upper Cased Name
const UpperCased = UpperCase(folderName)

//Final Path to the new Folder 
const finalPath = path.join(folderPath, UpperCased)


//Creating the folder
fs.mkdir(finalPath, (err) => {
    if (err) {
        setHasError(true)
        console.log(err)
    }
    else console.log("Folder Created Successfully")
})

if (hasError) throw new Error("Cannot Procceed ! Error Occured")



const ScriptTemplate = argv.template || template
if (!checkTemplate(ScriptTemplate, scriptPossible)) throw new Error("Template must be : typescript or javascript")
const mainExtension = ScriptTemplate === "typescript" ? "tsx" : "jsx"
const main = path.join(finalPath, `${UpperCased}.${mainExtension}`)


const styleType = argv.style || style
if (!checkTemplate(styleType, stylePossible)) throw new Error(`${styleType}Is not a valid style type [css,scss,sass]`)
const stylePath = path.join(finalPath, `${UpperCased}.${styleType}`)


if (hasError) throw new Error("Problem Occurred")



//Main File Template 
const fileTemplate = argv.type || type
if (!checkTemplate(fileTemplate, componentPossible)) throw new Error("File Template : class : Class Component || function : Functionnal Component")
//Choose the template according to 
let toModifyTemplate
fileTemplate === "function" ?
    toModifyTemplate = fs.readFileSync(__dirname + '/template/rfce.js').toString()
    :
    toModifyTemplate = fs.readFileSync(__dirname + '/template/rce.js').toString()

//Change the name of the template to match the same folder component name
const finalTemplate = toModifyTemplate.replaceAll("rfce", UpperCased).replaceAll("rce", UpperCased)
//Create and write to the file the template
fs.writeFile(main, finalTemplate, (err) => {
    if (err) throw new Error(err)
})
//Create the style file 
const styleFile = createFile(stylePath, styleType)
if (!styleFile) throw new Error("Cannot Procceed ! Error Occured")

const testPath = path.join(finalPath, `/${UpperCased}.test.js`)
fs.writeFile(testPath, "", (err) => {
    if (err) throw new Error(err)
})







