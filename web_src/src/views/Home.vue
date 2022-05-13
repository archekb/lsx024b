<template>
    <v-card flat class="mx-auto" max-width="1915">
        <v-toolbar color="primary" dark extended flat>
            <v-toolbar-title>LS{{ this.state?.controller_rated?.output_current_of_load || 'X0' }}24B Monitor</v-toolbar-title>
        </v-toolbar>

        <v-card class="mx-auto" max-width="1020" style="margin-top: -50px;">
            <v-toolbar flat>
                <v-toolbar-title class="grey--text"> {{ nameFormatter(tab) }} </v-toolbar-title>
                <v-btn icon @click="modern = !modern" :active="modern"><v-icon>mdi-apps</v-icon></v-btn> 
                <v-menu anchor="start">
                    <template v-slot:activator="{ props }">
                        <v-btn icon v-bind="props">
                            <v-icon>mdi-dots-vertical</v-icon>
                        </v-btn>
                    </template>

                    <v-list class="grey lighten-3">
                        <v-list-item v-for="item, key in items" :key="key" :disabled="key == tab" @click="tab = key">
                            {{ nameFormatter(key) }}
                        </v-list-item>
                    </v-list>
                </v-menu>
            </v-toolbar>

            <v-divider></v-divider>

            <v-window v-model="tab">
                <v-window-item v-for="item, key in items" :key="key" :value="key" style="min-height: 200px; margin: 10px;">
                    <v-card flat>
                        <ControllerRealTime v-if="key == 'controller_real_time' && modern" :data="item" :nameFormatter="nameFormatter" :unit="unit" />
                        <v-table v-else fixed-header>
                            <thead>
                                <tr>
                                    <th class="text-left">Name</th>
                                    <th class="text-left">Value</th>
                                </tr>
                            </thead>
                            <tbody>
                            <tr v-for="[k, v]  in Object.entries(item)" :key="k">
                                <td>{{ nameFormatter(k) }}</td>
                                <td>{{ v }} {{ unit(k, v) }}</td>
                            </tr>
                            </tbody>
                        </v-table>
                    </v-card>
                </v-window-item>
            </v-window>
        </v-card>
    </v-card>

</template>

<script>
import ControllerRealTime from "@/components/ControllerRealTime"
import { useStateStore } from "@/store/state";
import { storeToRefs } from 'pinia'
import { onUnmounted } from 'vue';

export default {
    name: 'home-page',
    components: {
        ControllerRealTime,
    },
    setup() {
        const stateStore = useStateStore();
        const { state, inProgress, error } = storeToRefs(stateStore);

        stateStore.start()
        onUnmounted(stateStore.stop)
        return {
            state,
            inProgress,
            error
        }
    },
    data() {
        return {
            tab: "controller_real_time",
            modern: true
        }
    },
    methods: {
        nameFormatter(str){
            return str.split('_').map(x => x[0].toUpperCase() + x.slice(1)).join(' ')
        },
        unit(name, value) {
            name = name.toLowerCase();
            if (typeof value === 'number' && isFinite(value)) {
                if (name.includes("temperature")) return "℃";
                if (name.includes("voltage")) return "V";
                if (name.includes("power")) return "W";
                if (name.includes("current")) return "A";
                if (name.includes("energy")) return "KWh"
                if (name.includes("soc")) return "%"
            }
            return ""
        }
    },
    computed: {
        items() {
            return Object.fromEntries(Object.entries(this.state).filter(([a])=> a.includes("controller")))
        }
    }
 }
</script>
