let vm = new Vue( {
    el: "#app",
    data: {
        doneList: [],
        undoList: [],
        item: ''
    },
    components: {
        'todo-item': {
            name: 'todo-item',
            template: `
            <div class="item">
                <input type="checkbox" :checked="mydone" :disabled="disabled" @click="handleClick">
                <span>{{text}}</span>
                <button :data-idx="idx"> x </button>
            </div>
            `,
            props: {
                done: {
                    type: Boolean,
                    default: false
                },
                disabled: {
                    type: Boolean,
                    default: false
                },
                text: {
                    type: String,
                    required: true
                },
                idx: {
                    type: Number,
                    required: true
                }
            },
            data () {
                return {
                    mydone: this.done
                }
            },
            methods: {
                handleClick () {
                    if ( this.mydone ) {
                        this.mydone = false;
                        return;
                    } else {
                        this.mydone = true
                        this.$emit( 'check', this.idx );
                    }
                }
            },
        },
        'todo-header': {
            name: 'todo-header',
            template: `
            <p class="title">
                <span>{{text}}</span><span class="num" id="udNum">{{num}}</span>
            </p>
            `,
            props: {
                text: {
                    type: String
                },
                num: {
                    type: Number,
                    default: 0
                }
            }
        }
    },
    methods: {
        handleSubmit () {
            this.undoList.unshift( {
                text: this.item,
                idx: new Date().getTime() % 1000000
            } );
            this.item = ''
        },
        handleCheck ( idx ) {
            console.log( `gotta ${ idx } ` + Math.random() );
            let temp;
            for ( const item in this.undoList ) {
                if ( this.undoList[ item ].idx == idx ) {
                    temp = this.undoList.splice( item, 1 )[ 0 ];
                    break;
                }
            }
            this.doneList.unshift( temp );
        }
    }
} )


var url = "http://api.cugxuan.cn/api/hackday/testquery?name=123345";
function push() {
    


    var xmlhttp =new XMLHttpRequest();

    //alert(xmlhttp.readyState);

    xmlhttp.onreadystatechange=function() {
        if(xmlhttp.readyState == 4 && xmlhttp.status == 200)

        //alert(xmlhttp.readyState);
        alert(xmlhttp.responseText);

        var jsonObj = JSON.parse(xmlhttp.responseText);//JSON.parse() returns JSON object 
        alert(jsonObj.name);
        //document.getElementById("date").innerHTML = jsonObj.date;  
        // document.getElementById("time").innerHTML = jsonObj.time;

    }
    xmlhttp.open("GET",url,true);
    xmlhttp.send();
}

/*
function push1() {

    //alert("123");
    $.ajax({
        type: "get",
        url: "http://api.cugxuan.cn/api/hackday/testquery?name=123345",
        dataType: 'jsonp', //【jsonp进行跨域请求 只支持get】
        //data:{ //【这里填写是传给服务端的数据 可传可不传 数据必须是json格式】
        //    "a":"b",
        //    "c":"d"
        //},
        success: function(result) { //【成功回调】
            alert(result);
        },
        error: function(xhr, type) { //【失败回调】
        }
        });
}
*/