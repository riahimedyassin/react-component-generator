const fs = require("fs")
const UpperCase = require("../utils/UppedCase")
const path = require("path")

class FileManager {
    createFile(path, data = "") {
        fs.writeFile(path, data, (err) => {
            if (err) console.log(err)
            process.exit(1)
        })
    }
    validateFolder(name) {
        if (!name) return null;
        if (name.indexOf("-") === 0) return null;
        //Parse to Valide Name
        return UpperCase(name.trim().split(" ").join("-"))
    }
    getFinalPath(dist, validName) {
        if (validName) {
            const folderPath = `./src/${dist}`;
            const finalPath = path.join(folderPath, validName);
            return finalPath
        }
        return null
    }
    createFolder(dist, name) {
        const validName = this.validateFolder(name)
        if (validName) {
            const finalPath = this.getFinalPath(dist, validName)
            fs.mkdir(finalPath, (err) => {
                if (err) {
                    console.log(err)
                    process.exit(1)
                }
            })
        }
        else {
            console.log("Error Creating the folder : Invalid Folder Name")
            process.exit(1)
        }
    }
    

}

module.exports = FileManager