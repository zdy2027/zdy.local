function uploadConf() {
    var mess = document.getElementById("middle-wizard");
    var inputs = mess.getElementsByTagName("input");
    var select = document.getElementsByTagName("select");
    var parameter = {};
    for (i=0;i<select.length;i++)
    {
        if (select[i].name=="os")
        {
            parameter[select[i].name] = select[i].value;
        }
        else if (select[i].name=="update_cycle")
        {
            parameter[select[i].name] = select[i].value;
        }
        else if (select[i].name=="upload_cycle")
        {
            parameter[select[i].name] = select[i].value;
        }
    }


    for(i=0;i<inputs.length;i++)
    {
        if (inputs[i].name=="upload")
        {
            parameter[inputs[i].name] = inputs[i].value;
        }
        else if (inputs[i].name=="preupload")
        {
            parameter[inputs[i].name] = inputs[i].value;
        }
        else if (inputs[i].name=="server")
        {
            parameter[inputs[i].name] = inputs[i].value;
        }
        else if (inputs[i].name=="lastmodify")
        {
            parameter[inputs[i].name] = inputs[i].value;
        }
        else if (inputs[i].name=="port")
        {
            parameter[inputs[i].name] = inputs[i].value;
        }
    }
    $.ajax({url:"/submit", type: "post", dataType: "json", data: JSON.stringify(parameter), contentType: "application/json", success:function(result){
        if (result["status"]=="succeed")
        {
            self.location='/';
        }else
        {
            self.location='/conf';
        }
    }});
}