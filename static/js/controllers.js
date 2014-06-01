var app = angular.module('bank', ['dateInput']);

app.controller('transactionsList', function($scope, $http) {
    $http.get('/transactions').success(function(data) {
        $scope.transactions=data;
    });

    $scope.endDate = new Date();
    $scope.startDate = new Date();
    $scope.startDate.setMonth($scope.startDate.getMonth() - 1);

    $scope.dateFilter = function(transaction) {
      var transDate = new Date(transaction.Date);
      return transDate >= $scope.startDate && transDate <= $scope.endDate;
    };
});

