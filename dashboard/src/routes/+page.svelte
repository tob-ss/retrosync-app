<script lang="ts">
  import Background from "../lib/components/Background.svelte";

  function beforeunload(event: BeforeUnloadEvent) {
    event.preventDefault();
    return event.returnValue = '';
  }


</script>




<Background />

<svelte:window on:beforeunload={beforeunload}/>


<style>


</style>