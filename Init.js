#!/usr/bin/env node
'use strict';
const fs = require("node:fs")


class Init {
    run() {
        const file = fs.createWriteStream("./generator.config.json", { encoding: "utf-8" })
        file.write('{"style":"-css","template":"-js","type":"function","dist":"components","test":true}')
        file.end()
    }
}

const intit = new Init();
intit.run()


module.exports = Init


