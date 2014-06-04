angular.module('bankFilters', [])
    .filter('utcDate', function() {
        return function(input) {
            var date = new Date(input);
            var month = date.getUTCMonth() < 10 ? "0" + date.getUTCMonth() : date.getUTCMonth();
            var day = date.getUTCDate() < 10 ? "0" + date.getUTCDate() : date.getUTCDate();
            return "" + date.getUTCFullYear() + "-" + month + "-" + day;
        }
    });
