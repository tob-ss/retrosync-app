<script lang="ts">
    import { Grid } from "@svar-ui/svelte-grid";
    import { GetSaves } from '../../wailsjs/go/main/App';
    

  //  let data: Array<Record<string, any>> = GetSaves();

    let loadGrid: boolean = $state(false);

    const columns = [
            {id: "ID", header: "Unique Save Number"},
            {id: "Game_Name", header: "Game Name"},
            {id: "Console", header: "Console"},
            {id: "Device", header: "Current Device"},
            {id: "Time_Modified", header: "Last Saved"},
            {id: "Save_Path", header: "Save Location"},
        ]

    GetSaves().then(data => {
        loadGrid = true;
    });

    let data = $state([]);

    (async () => {
        const data = await GetSaves();
        loadGrid = true;
    });
    
</script>


<div class="h-full w-full">
    <div class="h-full grid grid-cols-3 grid-rows-6 gap-4">
        <div class="col-span-3 row-span-6">
        {#if loadGrid === true}
          <Grid {data} {columns} />
        {/if}
        </div>
    </div>
</div>

<style>


</style>
