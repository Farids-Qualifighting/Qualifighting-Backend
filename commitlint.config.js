/** @type {import('@commitlint/types').UserConfig} */
const CommitLintConfiguration = {
    extends: ["@commitlint/config-conventional"],
    rules: {
        // add your own scope here if needed
        "scope-enum": [
            2,
            "always",
            ["competitions", "notes", "parents", "schools", "sports", "students", "teachers", "trainers", "config", "DO", "encryption", "firebase", "mongodb", "deps"],
        ],
        "scope-case": [2, "always", "kebab-case"],
    },
};

module.exports = CommitLintConfiguration;