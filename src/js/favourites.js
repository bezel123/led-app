$(document).ready(function () {
    
    /*
        variables
    */

    var name, color;

    /*
        functions to convert hex to rgb or vice versa
    */

    function byte2Hex(n) {
        var nybHexString = "0123456789ABCDEF";
        return (
            String(nybHexString.substr((n >> 4) & 0x0f, 1)) +
            nybHexString.substr(n & 0x0f, 1)
        );
    }

    function RGB2Color(r, g, b) {
        return "#" + byte2Hex(r) + byte2Hex(g) + byte2Hex(b);
    }

    function hexToRgb(hex) {
        // Expand shorthand form (e.g. "03F") to full form (e.g. "0033FF")
        var shorthandRegex = /^#?([a-f\d])([a-f\d])([a-f\d])$/i;
        hex = hex.replace(shorthandRegex, function (m, r, g, b) {
            return r + r + g + g + b + b;
        });

        var result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(hex);
        return result ? {
                r: parseInt(result[1], 16),
                g: parseInt(result[2], 16),
                b: parseInt(result[3], 16)
            } :
            null;
    }

    //sets rgb values to slider
    function setColor(r, g, b) {
        $("#slider-r-preview")[0].value = r;
        $("#slider-g-preview")[0].value = g;
        $("#slider-b-preview")[0].value = b;
    }

    /*
        event functions
    */

    //applys slider to preview
    $(".rgb-slider-preview").on("input", function () {
        var obj = {
            name: $("#name")[0].value,
            r: $("#slider-r-preview")[0].value,
            g: $("#slider-g-preview")[0].value,
            b: $("#slider-b-preview")[0].value
        };

        $(".preview-color").css("background-color", RGB2Color(obj.r, obj.g, obj.b));
        color = RGB2Color(obj.r, obj.g, obj.b);
        $("#color")[0].value = color
    });

    //applys input to preview and slider
    $("#color").keyup(function (event) {
        let input = $(this)[0].value;
        if (input.charAt(0) == "#") {
            try {
                $(".preview-color").css("background-color", input);
                setColor(hexToRgb(input).r, hexToRgb(input).g, hexToRgb(input).b)
                color = input;
                $("#color")[0].value = color;
            } catch (error) {}
        }
    });

    //send settings to rest
    $(".create").on("click", function () {
        name = $("#name")[0].value;
        if (name == "") {
            alert("Please enter a name")
        } else if (color == null) {
            alert("Please enter a color")
        }
        console.log("Alles Top");
    });
});