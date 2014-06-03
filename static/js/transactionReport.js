var app = angular.module('bank');

// A service for generating graphable accumulation reports from lists of credits & debits
app.factory('transReportFactory', function() {
    return {
        'foo': function() {
            console.log('bar');
        }
    }
});
