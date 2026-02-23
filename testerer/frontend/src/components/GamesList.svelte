<script lang="ts">
    import { GetSaves } from "../../wailsjs/go/main/App";
    import { GetHeaders } from "../../wailsjs/go/main/App";

  //  let data: Array<Record<string, any>> = GetSaves();

    GetSaves().then(data => {
        console.log(data);
    });

    let data = $state([]);

    (async () => {
        const data = await GetSaves();
        
    });

    let headers = $state([]);

    GetHeaders(data).then(headers => {
        console.log(headers);
    });

    (async () => {
        const headers = await GetHeaders(data);

        
    });
    


    
    
</script>


<div class="h-full w-full">
    
        {#await GetSaves()}
            <p>Getting Saves...</p>
        {:then data}
        {#await GetHeaders(data)}
            <p>Getting Headers...</p>
        {:then headers}
        <div class="h-full grid grid-cols-3 grid-rows-6 gap-4 text-[#D7D6FC] font-heebo">
        {#each headers as header}
            <div class="col-span-3">
            <h1>{header}</h1>
                <div class="col-span-3 row-span-5 row-start-2">
                {#each data as save}
                {#if header === save.Date_String}
                <p>Showing game: {save.Game_Name}</p>
                {/if}
                {/each}
                </div>
            </div>
        {/each}
        </div>
        {/await}
        {/await}
    
</div>