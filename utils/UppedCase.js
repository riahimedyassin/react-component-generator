/**
 *
 * @param {string} str
 * @throws {Error} If the Argument is not a string or there is no argument , throws an error
 * @returns {string}
 *
 */
const UpperCase = (str) => {
    if (!str) throw new Error("Parameter : Require a String ")
    return str.charAt(0).toUpperCase() + str.slice(1)
}

module.exports = UpperCase