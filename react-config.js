#!/usr/bin/env node
"use strict";

const fs = require("node:fs")

const file = fs.createWriteStream("./generator.config.json",{encoding:"utf-8"})
file.write('{"style":"css","template":"javascript","type":"f","dist":"components"}')
file.end()

