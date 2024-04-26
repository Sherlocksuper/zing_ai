const Mock = require('mockjs');
Mock.mock('/api/user', 'get', {
    code: 200,
    'data|10-20': [{
        'id|+1': 1,
        'name': '@cname',
        'age|18-28': 1,
        'address': '@county(true)',
        'avatar': '@image("200x100", "#50B347", "#FFF", "Mock.js")',
        'time': '@datetime'
    }]
});