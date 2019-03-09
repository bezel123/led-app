$(document).ready(function () {
    $('.rgb-slider').on('input', function () {
        //console.log(this.value);
        //console.log(this.id)
        var r = $('#slider-r')[0].value
        var g = $('#slider-g')[0].value
        var b = $('#slider-b')[0].value
        console.log(r, g, b);

    });
});