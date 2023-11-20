#!/usr/bin/env node
'use strict';
/**
 * @requires fs
 * @requires path
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
const checkTemplate = require("./utils/checkTemplate");



const FileManager = require("./Manager/FileManager")
const ExtensionManager = require("./Manager/ExtensionManager")
const { scriptPossible, stylePossible, componentPossible } = require("./utils/constants")
const Init = require("./Init")
let configFile = fs.readFileSync("./generator.config.json", { encoding: "utf-8" }).toString()
let generator ;
if(!configFile)  {
    const init = new Init()
    init.run()
}
configFile = fs.readFileSync("./generator.config.json", { encoding: "utf-8" }).toString()
generator=JSON.parse(configFile)

const { style, template, type, dist, test } = generator

class Main {
    file
    extension
    name
    constructor() {
        this.file = new FileManager()
        this.extension = new ExtensionManager()
        this.name = process.argv.slice(2, 3)[0]
    }
    getDist() {
        return dist
    }
    getValideFolderName(name) {
        return this.file.validateFolder(name)
    }
    createFolder(dist, validName) {
        this.file.createFolder(dist, validName)
    }
    getFinalPath(dist, validName) {
        return this.file.getFinalPath(dist, validName)
    }

    getScriptPath(finalPath, valideFolderName) {
        const scriptExtension = this.extension.mainValidation(template, "script", scriptPossible)
        const mainScript = path.join(finalPath, `${valideFolderName}.${scriptExtension}`)
        return mainScript
    }

    getTemplate(valideFolderName) {
        const fileTemplate = type
        if (!checkTemplate(fileTemplate, componentPossible)) throw new Error("File Template : class : Class Component || function : Functionnal Component : Specify that in the generator.config file ")
        const toModifyTemplate = fileTemplate === "function" ? fs.readFileSync(__dirname + '/template/rfce.js').toString() : fs.readFileSync(__dirname + '/template/rce.js').toString()
        const finalTemplate = toModifyTemplate.replaceAll("rfce", valideFolderName).replaceAll("rce", valideFolderName)
        return finalTemplate
    }

    getStylePath(finalPath, valideFolderName) {
        const styleExtension = this.extension.mainValidation(style, "style", stylePossible)
        const stylePath = path.join(finalPath, `${valideFolderName}.${styleExtension}`)
        return stylePath
    }

    createFile(filePath, content = "") {
        this.file.createFile(filePath, content)
    }

    main() {
        const name = this.name
        const dist = this.getDist()

        const valideFolder = this.getValideFolderName(name)
        if (!valideFolder) {
            console.log("Invalid Folder Name")
            process.exit(1)
        }
        const finalPath = this.getFinalPath(dist, valideFolder)

        const mainScriptPath = this.getScriptPath(finalPath, valideFolder);
        const finalTemplate = this.getTemplate(valideFolder)
        const mainStylePath = this.getStylePath(finalPath, valideFolder)




        this.createFolder(dist, name)
        console.log("Folder Created Successfully")
        this.createFile(mainScriptPath, finalTemplate)
        console.log("Script File has been created Successfully")
        this.createFile(mainStylePath)
        console.log("Style File has been created Successfully")
        if (test) {
            const testPath = path.join(finalPath, `${valideFolder}.test.js`)
            this.createFile(testPath)
            console.log("Test File has been created Successfully")
        }

        process.exit(0)

    }


}

const app = new Main();
app.main()