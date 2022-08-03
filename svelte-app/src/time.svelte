<script>
    import _ from "lodash";

    export let time;
    let timeText = time;
    let saveInput = '12:15'

    function formatTime(e) {

        if(e.key == "Delete" || e.key == "Backspace"){
            return
        }
        let v = e.target.value;

        if (v.length === 2 && v[2] !== ":") {
            e.target.value = v + ":";
        }

        if(v.length === 6){
            e.target.value =saveInput
            return
        }

        saveInput = e.target.value





        time = _.replace(timeText, ":", "");
    }
</script>

<input
    type="text"
    placeholder="Time"
    value="12:15"
    class="pl-2 p-4 mb-2"

    on:keyup={formatTime}
    id="time"
/>
