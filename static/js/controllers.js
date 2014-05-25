var app = angular.module('bank', []);

app.controller('transactionsList', function($scope, $http) {
    $http.get('/transactions').success(function(data) {
        $scope.transactions=data;
    });
});
