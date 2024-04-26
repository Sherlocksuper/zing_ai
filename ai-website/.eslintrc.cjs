module.exports = {
    root: true,
    env: {browser: true, es2020: true, jest: true, node: true,},
    extends: [
        'eslint:recommended',
        'plugin:@typescript-eslint/recommended',
        'plugin:react-hooks/recommended',
    ],
    ignorePatterns: ['dist', '.eslintrc.cjs'],
    parser: '@typescript-eslint/parser',
    plugins: ['react-refresh'],
    rules: {
        'no-constant-condition': ['error', {checkLoops: false}],
        '@typescript-eslint/no-explicit-any': 'off',
        "@typescript-eslint/ban-ts-comment": "off",
        'react-refresh/only-export-components': [
            'warn',
            {allowConstantExport: true},
        ],
    },
    overrides: [
        {
            files: ['test/**/*.ts', 'test/**/*.tsx'],
            rules: {
                '@typescript-eslint/no-var-requires': 'off',
            },
            env: {jest: true},
        },
    ],
}
