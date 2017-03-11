// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package define

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type MerchantService interface {
	// Parameters:
	//  - MchId
	Complex(mchId int32) (r *ComplexMerchant, err error)
	// Parameters:
	//  - Usr
	//  - OriPwd
	CheckLogin(usr string, oriPwd string) (r *Result_, err error)
	// Parameters:
	//  - MchId
	Stat(mchId int32) (r *Result_, err error)
}

type MerchantServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewMerchantServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *MerchantServiceClient {
	return &MerchantServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewMerchantServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *MerchantServiceClient {
	return &MerchantServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - MchId
func (p *MerchantServiceClient) Complex(mchId int32) (r *ComplexMerchant, err error) {
	if err = p.sendComplex(mchId); err != nil {
		return
	}
	return p.recvComplex()
}

func (p *MerchantServiceClient) sendComplex(mchId int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("Complex", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := MerchantServiceComplexArgs{
		MchId: mchId,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *MerchantServiceClient) recvComplex() (value *ComplexMerchant, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "Complex" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "Complex failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "Complex failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error7 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error8 error
		error8, err = error7.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error8
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "Complex failed: invalid message type")
		return
	}
	result := MerchantServiceComplexResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Usr
//  - OriPwd
func (p *MerchantServiceClient) CheckLogin(usr string, oriPwd string) (r *Result_, err error) {
	if err = p.sendCheckLogin(usr, oriPwd); err != nil {
		return
	}
	return p.recvCheckLogin()
}

func (p *MerchantServiceClient) sendCheckLogin(usr string, oriPwd string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("CheckLogin", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := MerchantServiceCheckLoginArgs{
		Usr:    usr,
		OriPwd: oriPwd,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *MerchantServiceClient) recvCheckLogin() (value *Result_, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "CheckLogin" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "CheckLogin failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "CheckLogin failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error9 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error10 error
		error10, err = error9.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error10
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "CheckLogin failed: invalid message type")
		return
	}
	result := MerchantServiceCheckLoginResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - MchId
func (p *MerchantServiceClient) Stat(mchId int32) (r *Result_, err error) {
	if err = p.sendStat(mchId); err != nil {
		return
	}
	return p.recvStat()
}

func (p *MerchantServiceClient) sendStat(mchId int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("Stat", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := MerchantServiceStatArgs{
		MchId: mchId,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *MerchantServiceClient) recvStat() (value *Result_, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "Stat" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "Stat failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "Stat failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error11 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error12 error
		error12, err = error11.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error12
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "Stat failed: invalid message type")
		return
	}
	result := MerchantServiceStatResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type MerchantServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      MerchantService
}

func (p *MerchantServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *MerchantServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *MerchantServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewMerchantServiceProcessor(handler MerchantService) *MerchantServiceProcessor {

	self13 := &MerchantServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self13.processorMap["Complex"] = &merchantServiceProcessorComplex{handler: handler}
	self13.processorMap["CheckLogin"] = &merchantServiceProcessorCheckLogin{handler: handler}
	self13.processorMap["Stat"] = &merchantServiceProcessorStat{handler: handler}
	return self13
}

func (p *MerchantServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x14 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x14.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x14

}

type merchantServiceProcessorComplex struct {
	handler MerchantService
}

func (p *merchantServiceProcessorComplex) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := MerchantServiceComplexArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Complex", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := MerchantServiceComplexResult{}
	var retval *ComplexMerchant
	var err2 error
	if retval, err2 = p.handler.Complex(args.MchId); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Complex: "+err2.Error())
		oprot.WriteMessageBegin("Complex", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("Complex", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type merchantServiceProcessorCheckLogin struct {
	handler MerchantService
}

func (p *merchantServiceProcessorCheckLogin) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := MerchantServiceCheckLoginArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("CheckLogin", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := MerchantServiceCheckLoginResult{}
	var retval *Result_
	var err2 error
	if retval, err2 = p.handler.CheckLogin(args.Usr, args.OriPwd); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing CheckLogin: "+err2.Error())
		oprot.WriteMessageBegin("CheckLogin", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("CheckLogin", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type merchantServiceProcessorStat struct {
	handler MerchantService
}

func (p *merchantServiceProcessorStat) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := MerchantServiceStatArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Stat", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := MerchantServiceStatResult{}
	var retval *Result_
	var err2 error
	if retval, err2 = p.handler.Stat(args.MchId); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Stat: "+err2.Error())
		oprot.WriteMessageBegin("Stat", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("Stat", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - MchId
type MerchantServiceComplexArgs struct {
	MchId int32 `thrift:"mchId,1" json:"mchId"`
}

func NewMerchantServiceComplexArgs() *MerchantServiceComplexArgs {
	return &MerchantServiceComplexArgs{}
}

func (p *MerchantServiceComplexArgs) GetMchId() int32 {
	return p.MchId
}
func (p *MerchantServiceComplexArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *MerchantServiceComplexArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.MchId = v
	}
	return nil
}

func (p *MerchantServiceComplexArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Complex_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *MerchantServiceComplexArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("mchId", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:mchId: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.MchId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.mchId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:mchId: ", p), err)
	}
	return err
}

func (p *MerchantServiceComplexArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MerchantServiceComplexArgs(%+v)", *p)
}

// Attributes:
//  - Success
type MerchantServiceComplexResult struct {
	Success *ComplexMerchant `thrift:"success,0" json:"success,omitempty"`
}

func NewMerchantServiceComplexResult() *MerchantServiceComplexResult {
	return &MerchantServiceComplexResult{}
}

var MerchantServiceComplexResult_Success_DEFAULT *ComplexMerchant

func (p *MerchantServiceComplexResult) GetSuccess() *ComplexMerchant {
	if !p.IsSetSuccess() {
		return MerchantServiceComplexResult_Success_DEFAULT
	}
	return p.Success
}
func (p *MerchantServiceComplexResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MerchantServiceComplexResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *MerchantServiceComplexResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &ComplexMerchant{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *MerchantServiceComplexResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Complex_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *MerchantServiceComplexResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *MerchantServiceComplexResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MerchantServiceComplexResult(%+v)", *p)
}

// Attributes:
//  - Usr
//  - OriPwd
type MerchantServiceCheckLoginArgs struct {
	Usr    string `thrift:"usr,1" json:"usr"`
	OriPwd string `thrift:"oriPwd,2" json:"oriPwd"`
}

func NewMerchantServiceCheckLoginArgs() *MerchantServiceCheckLoginArgs {
	return &MerchantServiceCheckLoginArgs{}
}

func (p *MerchantServiceCheckLoginArgs) GetUsr() string {
	return p.Usr
}

func (p *MerchantServiceCheckLoginArgs) GetOriPwd() string {
	return p.OriPwd
}
func (p *MerchantServiceCheckLoginArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *MerchantServiceCheckLoginArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Usr = v
	}
	return nil
}

func (p *MerchantServiceCheckLoginArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.OriPwd = v
	}
	return nil
}

func (p *MerchantServiceCheckLoginArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("CheckLogin_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *MerchantServiceCheckLoginArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("usr", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:usr: ", p), err)
	}
	if err := oprot.WriteString(string(p.Usr)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.usr (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:usr: ", p), err)
	}
	return err
}

func (p *MerchantServiceCheckLoginArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("oriPwd", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:oriPwd: ", p), err)
	}
	if err := oprot.WriteString(string(p.OriPwd)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.oriPwd (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:oriPwd: ", p), err)
	}
	return err
}

func (p *MerchantServiceCheckLoginArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MerchantServiceCheckLoginArgs(%+v)", *p)
}

// Attributes:
//  - Success
type MerchantServiceCheckLoginResult struct {
	Success *Result_ `thrift:"success,0" json:"success,omitempty"`
}

func NewMerchantServiceCheckLoginResult() *MerchantServiceCheckLoginResult {
	return &MerchantServiceCheckLoginResult{}
}

var MerchantServiceCheckLoginResult_Success_DEFAULT *Result_

func (p *MerchantServiceCheckLoginResult) GetSuccess() *Result_ {
	if !p.IsSetSuccess() {
		return MerchantServiceCheckLoginResult_Success_DEFAULT
	}
	return p.Success
}
func (p *MerchantServiceCheckLoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MerchantServiceCheckLoginResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *MerchantServiceCheckLoginResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &Result_{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *MerchantServiceCheckLoginResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("CheckLogin_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *MerchantServiceCheckLoginResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *MerchantServiceCheckLoginResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MerchantServiceCheckLoginResult(%+v)", *p)
}

// Attributes:
//  - MchId
type MerchantServiceStatArgs struct {
	MchId int32 `thrift:"mchId,1" json:"mchId"`
}

func NewMerchantServiceStatArgs() *MerchantServiceStatArgs {
	return &MerchantServiceStatArgs{}
}

func (p *MerchantServiceStatArgs) GetMchId() int32 {
	return p.MchId
}
func (p *MerchantServiceStatArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *MerchantServiceStatArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.MchId = v
	}
	return nil
}

func (p *MerchantServiceStatArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Stat_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *MerchantServiceStatArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("mchId", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:mchId: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.MchId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.mchId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:mchId: ", p), err)
	}
	return err
}

func (p *MerchantServiceStatArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MerchantServiceStatArgs(%+v)", *p)
}

// Attributes:
//  - Success
type MerchantServiceStatResult struct {
	Success *Result_ `thrift:"success,0" json:"success,omitempty"`
}

func NewMerchantServiceStatResult() *MerchantServiceStatResult {
	return &MerchantServiceStatResult{}
}

var MerchantServiceStatResult_Success_DEFAULT *Result_

func (p *MerchantServiceStatResult) GetSuccess() *Result_ {
	if !p.IsSetSuccess() {
		return MerchantServiceStatResult_Success_DEFAULT
	}
	return p.Success
}
func (p *MerchantServiceStatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MerchantServiceStatResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *MerchantServiceStatResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &Result_{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *MerchantServiceStatResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Stat_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *MerchantServiceStatResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *MerchantServiceStatResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MerchantServiceStatResult(%+v)", *p)
}
