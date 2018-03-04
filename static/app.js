(function(angular) {
    var app = angular.module('App', []);
    app.service("duaListBoxService", function() {
        var service = this;
        service.duallistbox_ins = {};

        // create a class
        var dualListBoxIns = function(){};
        dualListBoxIns.prototype.leftSelect = {};
        dualListBoxIns.prototype.rightSelect = {};

        dualListBoxIns.prototype.setLeftSelect = {};
        dualListBoxIns.prototype.getLeftSelect = {};
        dualListBoxIns.prototype.setRightSelect = {};
        dualListBoxIns.prototype.getRightSelect = {};

        // create new instance
        service.get_new_duallistbox = function(ins_name) {
            service.duallistbox_ins[ins_name] = new dualListBoxIns();
            return service.duallistbox_ins[ins_name];
        };

        // get instance
        service.get_duallistbox = function(ins_name) {
            return service.duallistbox_ins[ins_name];
        };

        // get right select
        this.getRightSelect = function(ins_name) {
            return service.duallistbox_ins[ins_name].rightSelect;
        };

        // set right select
        this.setRightSelect = function(ins_name, data) {
            service.duallistbox_ins[ins_name].rightSelect = data;
        };

        // get left select
        this.getLeftSelect = function(ins_name) {
            return service.duallistbox_ins[ins_name].leftSelect;
        };

        // set left select
        this.setLeftSelect = function(ins_name, data) {
            service.duallistbox_ins[ins_name].leftSelect = data;
        };
    });

    app.component('duallistbox', {
        templateUrl: '/duallistbox.html',
    });

    function duallistboxComponentController($scope, dualListBoxService) {
        var cur_ctl = this;
        this.$onInit = function() {
            this.dualListBoxService = dualListBoxService;
            this.duallistbox_ins = duaListBoxService.get_duallistbox(cur_ctl.insName);
            console.log(cur_ctl.insName);

            this.leftSelectShow = _.cloneDeep(cur_ctl.duallistbox_ins.leftSelect);
            this.rightSelectShow = _.cloneDeep(cur_ctl.duallistbox_ins.rightSelect);
            this.leftPage = "1";
            this.rightPage = "1";
            this.loading = false;
        };

        this.changeLeftPage = function(text, page, pageSize, total) {
            cur_ctl.leftPage = page;
        };

        this.changeRightPage = function(text, page, pageSize, total) {
            cur_ctl.rightPage = page;
        };

        this.choosed = function(item) {
            //remove total left to right
            var removed = _.remove(cur_ctl.duallistbox_ins.leftSelect, function(cur_item) {
                return cur_item.value === item.value && !cur_item.disable;
            });

            cur_ctl.duallistbox_ins.rightSelect = _.uniqBy(_.concat(cur_ctl.duallistbox_ins.rightSelect, removed), 'value');

            var removed2 = _.remove(cur_ctl.leftSelectShow, function(cur_item) {
                return cur_item.value === item.value && !cur_item.disable;
            });

            cur_ctl.duallistbox_ins _.uniqBy(_.concat(cur_ctl.rightSelectShow))
        }
    }
})(angular);