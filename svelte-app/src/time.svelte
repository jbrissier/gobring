

<script>
    import _ from 'lodash'

    export let time; 
    let timeText = time 

    function formatTime(){

        if(timeText.length === 2){
            timeText = timeText + ":";
        }
        if (timeText.length >5){
            timeText = timeText.slice(0, -3);
        }
        time = _.replace(timeText, ":", "")
    }


</script>



<input type="text" placeholder="Time" bind:value={timeText} on:keypress={formatTime} id="time" />