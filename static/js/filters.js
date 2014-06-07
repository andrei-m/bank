var bankFilters = angular.module('bankFilters', []);
bankFilters.filter('utcDate', function($filter) {
    return function(input, format) {
        var date = new Date(input);
        date.setMinutes(date.getMinutes() + date.getTimezoneOffset());
        // 'date' renders in local time by default. The TZ offset is added
        // to convert 'local' to 'UTC' before rendering
        return $filter('date')(date, format);
    }
});
// convert an integer representing the number of cents to decimal (/ 100)
bankFilters.filter('toDecimal', function() {
    var toFixed = function (value, precision) {
        var power = Math.pow(10, precision || 0);
        return Math.round(value * power) / power;
    }
    return function(input) {
        return toFixed(input / 100, 2);
    }
});
// convert a decimal representation to integer (* 100)
bankFilters.filter('fromDecimal', function() {
    return function(input) {
        return input * 100;
    }
});
// render the decimal to two decimal places with a dollar sign
bankFilters.filter('currency', function() {
    return function(input) {
        return "$" + input.toFixed(2);
    }
});
