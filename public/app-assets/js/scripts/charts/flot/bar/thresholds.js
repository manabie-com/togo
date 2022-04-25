/*=========================================================================================
    File Name: thresholds.js
    Description: Flot thresholds chart
    ----------------------------------------------------------------------------------------
    Item Name: Stack - Responsive Admin Theme
    Version: 1.0
    Author: GeeksLabs
    Author URL: http://www.themeforest.net/user/geekslabs
==========================================================================================*/

// Thresholds chart
// ------------------------------
$(window).on("load", function(){

    var d1 = [];
    for (var i = 0; i <= 60; i += 1) {
        d1.push([i, parseInt(Math.random() * 30 - 10)]);
    }

    function plotWithOptions(t) {
        $.plot("#thresholds", [{
            data: d1,
            color: "#00A5A8",
            threshold: {
                below: t,
                color: "#F25E75"
            },
            lines: {
                steps: true
            }
        }],{
            grid: {
                borderWidth: 1,
                borderColor: "#e9e9e9",
                color: '#999',
                minBorderMargin: 20,
                labelMargin: 10,
                margin: {
                    top: 8,
                    bottom: 20,
                    left: 20
                },
            },
        });
    }

    plotWithOptions(0);
});