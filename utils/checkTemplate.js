
/**
 @param {string} value
 @param {string[]} possibleValues
 @returns boolean
 */
const checkTemplate = (value, possibleValues) => {

    return possibleValues.includes(value)
};
/**
 * @exports checkTemplate
 * */
module.exports=checkTemplate