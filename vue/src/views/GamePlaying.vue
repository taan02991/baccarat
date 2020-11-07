<template>
    <div class='container'>
        <div class='row pt-3'>
            <div class='col'>
                Game ID: 12
                <br>
                Betting time: 13s
            </div>
            <div class='col text-center'>
                <button type="button" class="btn btn-danger">Quit Room</button>
                <br>
                <br>
                Round: 1/8
            </div>
            <div class='col text-right'>
                Username: TestUser
                <br>
                Token: 500
            </div>
        </div>
        <div class='card-table row my-3'>
            <div class='col'>
                <Card :rank='r' :suit='s'></Card>
            </div>
            <div class='col'>
                <Card :rank='r' :suit='s'></Card>
            </div>
            <div class='col'>
                <Card v-for='i in 10' v-bind:key="i" :style="'position:absolute; left:' + (15 + i*1.5) + 'px; z-index: ' + i"></Card>
            </div>
            <div class='col'>
                <Card :rank='r' :suit='s'></Card>
            </div>
            <div class='col'>
                <Card :rank='r' :suit='s'></Card>
            </div>
        </div>
        <div class='row my-5'>
            <div class='col-5'>
                <UserList/>
            </div>
            <div class='col-3'>
                <div class="input-group mx-auto my-auto">
                    <div class="input-group-prepend">
                        <span class="input-group-text">Amount</span>
                    </div>
                    <input type="text" class="form-control" aria-label="Amount (to the nearest dollar)">
                </div>

            </div>
            <div class='col-4'>
                <button type="button" class="btn btn-lg btn-block btn-success">Player</button>
                <button type="button" class="btn btn-lg btn-block btn-danger">Banker</button>
                <button type="button" class="btn btn-lg btn-block btn-info">Tie</button>
            </div>
        </div>
        <p>Test Animation</p>
        <button @click='r=""; s=""'>close card</button>
        <button @click='r="1"; s="S"'>open card</button>
        <button type="button" @click='showModal'>Launch demo modal</button>

        <!-- Modal -->
        <div class="modal fade" id="modal" tabindex="-1">
        <div class="modal-dialog">
            <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title">Result</h5>
            </div>
            <div class="modal-body">
                <p>Winner: Banker</p>
                <p>Your bet: Banker</p>
                <p>Amount: 100</p>
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
            </div>
            </div>
        </div>
        </div>
    </div>
</template>
 
<script>
import Card from "@/components/Card.vue";
import UserList from "@/components/UserList.vue";

export default{
    components: {
        Card,
        UserList
    },
    data: function() {
        return {
            r: '1',
            s: 'S'
        }
    },
    methods: {
        showModal() {
            $("#modal").modal('toggle');
        }
    }
}


</script> 

<style scoped>

.card-table {
    background-color: green;
    border-radius: 10px;
    border: solid;
    border-color: white;
    padding: 20px;
}

</style>