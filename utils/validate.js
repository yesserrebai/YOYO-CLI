export function isValidString(input) {
    // Define a regular expression pattern
    // ^ - Start of the string
    // [a-zA-Z] - Match any letter (uppercase or lowercase)
    // - Match hyphen
    // _ Match underscore
    // $ - End of the string
    const pattern = /^[a-zA-Z\-_]+$/;
    // Test the input string against the pattern
    return pattern.test(input);
}