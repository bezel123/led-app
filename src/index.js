$(document).ready(function () {
    //react to changes from the slider
    // $('.rgb-slider').on('input', function () {
    //     //var obj = {
    //     //    name: $(".card-title").html(),
    //     //    r: $('#slider-r')[0].value,
    //     //    g: $('#slider-g')[0].value,
    //     //    b: $('#slider-b')[0].value
    //     //}

    //     $("#name").val(
    //         $(this)
    //         .closest(".card-content")
    //         .find(".card-title")[0].innerText
    //     );
    //     console.log($('#rgb-form').serialize());

    //     var url = "http://rest/v1/sendValues/"

    //     $.get("http://rest/v1/CheckConnection",
    //         function (data) {
    //             console.log(data);
    //             alert("Can not connect to device")
    //         },
    //         "json"
    //     );

    //     $.ajax({
    //         type: "POST",
    //         url: url,
    //         crossDomain: true,
    //         data: $('#rgb-form').serialize(),
    //         dataType: "text",
    //         success: function (response) {
    //             console.log(response);

    //         }
    //     });

    // });

    $(".switch").change(function (e) {
        e.preventDefault();
        if (!$(this).is(":checked")) {
            console.log("ayy");

        }
    });
});