import{T as B}from"./Title.3ec1fd38.js";import{_ as j}from"./TransactionTable.bce31510.js";import{e as i,r as n,o as t,g as p,w as m,c as _,v as C,u as s,F as b,h,y as v,t as x,x as F,a as r,d as T}from"./vendor.b3082d65.js";import{d as D}from"./useMock.aee5a14d.js";import"./index.53d3be81.js";import"./shuffle.cbddf0eb.js";const f=[{label:"\u4F59\u989D",key:"balance"},{label:"\u4EA4\u6613\u6570",key:"transactionCount"},{label:"\u603B\u652F\u51FA",key:"totalSent"},{label:"\u603B\u6536\u5165",key:"totalReceived"},{label:"\u603B\u624B\u7EED\u8D39",key:"totalFees"},{label:"\u8F6C\u8D26\u6B21\u6570",key:"tokenTransferCount"}],E=i({props:{data:{type:Object},maps:{type:Object,required:!0}},setup(e){return(k,u)=>{const a=n("a-descriptions-item"),c=n("a-descriptions");return e.data?(t(),p(c,{key:0,bordered:""},{default:m(()=>[(t(!0),_(b,null,C(s(f),o=>(t(),p(a,{label:o.label},{default:m(()=>[v(x(e.data[o.key]),1)]),_:2},1032,["label"]))),256))]),_:1})):h("",!0)}}}),g=T("h2",null,"\u4EA4\u6613\u5217\u8868",-1),q=i({props:{address:String},setup(e){const u=F({id:e.address||"123"}),{data:a,error:c}=D(u);return(o,A)=>{var d,l;const y=n("a-divider");return t(),_(b,null,[r(B,{title:"\u5730\u5740\u8BE6\u60C5"}),r(E,{data:(d=s(a))==null?void 0:d.summary,maps:s(f)},null,8,["data","maps"]),r(y),g,r(j,{data:(l=s(a))==null?void 0:l.transactions},null,8,["data"])],64)}}});export{q as default};