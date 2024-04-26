/**
 * For a detailed explanation regarding each configuration property, visit:
 * https://jestjs.io/docs/configuration
 */
/** @type {import('ts-jest/dist/types').InitialOptionsTsJest} */

module.exports = {
    preset: 'ts-jest',
    testEnvironment: 'node',
    collectCoverage: true, // 输出测试覆盖率报告
    //timeout
    testTimeout: 20000,
};
