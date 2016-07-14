function showFitnessData() {
	$.get("/data/", function(data) {
        loadData(data);
    });
}

/* Example data:
{
    "startDate":"2016-07-01T10:54:35.565852368+01:00",
    "days":7,
    "steps":[
        1026,
        5146,
        4264,
        4896,
        7923,
        10378,
        5874
    ],
    "weight":[
        83,
        83,
        83
    ],
    "height":null,
    "caloriesExpendedOnActivities":[
        1629.9945068359375,
        1773.43798828125,
        2031.5050048828125,
        1995.6978759765625,
        2105.984619140625,
        2403.6982421875,
        1983.80810546875
    ],
    "caloriesExpendedOnBMR":[
        327.47686767578125,
        105.46992492675781,
        75.08860778808594,
        47.90279006958008,
        14.623250007629395,
        35.55097579956055,
        40.21820068359375
    ]
}
*/

function loadData(d) {
    var myChart = new Chart(document.getElementById("myChart"), {
        type: 'line',
        data: {
            labels: d.dayNames,
            datasets: [{
                label: 'Steps',
                data: d.steps,
                backgroundColor: [
                    'rgba(255, 99, 132, 0.2)',
                    'rgba(54, 162, 235, 0.2)',
                    'rgba(255, 206, 86, 0.2)',
                    'rgba(75, 192, 192, 0.2)',
                    'rgba(153, 102, 255, 0.2)',
                    'rgba(255, 159, 64, 0.2)'
                ],
                borderColor: [
                    'rgba(255,99,132,1)',
                    'rgba(54, 162, 235, 1)',
                    'rgba(255, 206, 86, 1)',
                    'rgba(75, 192, 192, 1)',
                    'rgba(153, 102, 255, 1)',
                    'rgba(255, 159, 64, 1)'
                ],
                borderWidth: 1
            }]
        },
        options: {
            scales: {
                yAxes: [{
                    ticks: {
                        beginAtZero:true
                    }
                }]
            }
        }
    });

    var myChart = new Chart(document.getElementById("caloriesChart"), {
        type: 'bar',
        data: {
            labels: d.dayNames,
            datasets: [{
                label: 'Calories Expended on Activities',
                data: d.caloriesExpendedOnActivities,
                borderWidth: 1
            }]
        },
        options: {
            scales: {
                yAxes: [{
                    ticks: {
                        beginAtZero:true
                    }
                }]
            }
        }
    });
}