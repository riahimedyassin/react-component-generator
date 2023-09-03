



class ExtensionManager {
    args = process.argv;
    findParameter(data, values) {
        let i = 0;
        while (i < values.length && data.indexOf(values[i]) === -1) i++
        return data.indexOf(values[i]);

    }
    ValideExtension(extension, type) {
        const valide = extension.split("").slice(1,).join("")
        switch (type) {
            case "style": {
                return valide;
            }
            case "script": {
                return valide + "x"
            }
            default: return null
        }
    }

    mainValidation(def, type, possible) {
        const exist = this.findParameter(this.args, possible)
        if (exist != -1) {
            return this.ValideExtension(this.args[exist], type)
        }
        else {
            return this.ValideExtension(def, type)
        }
    }
}


module.exports = ExtensionManager

