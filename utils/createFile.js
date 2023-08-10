const fs = require("fs")

/**
 *  @async
 *  @function createFile
 *  @param {String} path Path to the file
 *  @param {String} extension File Extension
 *  @param {String} [data=""] Data To Append in the File
 *  @returns boolean
 */
const createFile = async (path, extension, data = "") => {
    fs.createWriteStream(path, {encoding: "utf-8"})
        .write(data, (error) => {
            if (error) {
                console.log(error)
                return false
            }
            console.log(`${extension} Created Successfully`)
            return true
        })
};

module.exports = createFile