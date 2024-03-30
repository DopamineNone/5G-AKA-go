// Code generated by thriftgo (0.3.6). DO NOT EDIT.

package udm

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type ProtocolService interface {
	HandleConnection(ctx context.Context, data string) (err error)
}

type ProtocolServiceClient struct {
	c thrift.TClient
}

func NewProtocolServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ProtocolServiceClient {
	return &ProtocolServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewProtocolServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ProtocolServiceClient {
	return &ProtocolServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewProtocolServiceClient(c thrift.TClient) *ProtocolServiceClient {
	return &ProtocolServiceClient{
		c: c,
	}
}

func (p *ProtocolServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *ProtocolServiceClient) HandleConnection(ctx context.Context, data string) (err error) {
	var _args ProtocolServiceHandleConnectionArgs
	_args.Data = data
	var _result ProtocolServiceHandleConnectionResult
	if err = p.Client_().Call(ctx, "HandleConnection", &_args, &_result); err != nil {
		return
	}
	return nil
}

type ProtocolServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      ProtocolService
}

func (p *ProtocolServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *ProtocolServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *ProtocolServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewProtocolServiceProcessor(handler ProtocolService) *ProtocolServiceProcessor {
	self := &ProtocolServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("HandleConnection", &protocolServiceProcessorHandleConnection{handler: handler})
	return self
}
func (p *ProtocolServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type protocolServiceProcessorHandleConnection struct {
	handler ProtocolService
}

func (p *protocolServiceProcessorHandleConnection) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ProtocolServiceHandleConnectionArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("HandleConnection", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := ProtocolServiceHandleConnectionResult{}
	if err2 = p.handler.HandleConnection(ctx, args.Data); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing HandleConnection: "+err2.Error())
		oprot.WriteMessageBegin("HandleConnection", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	}
	if err2 = oprot.WriteMessageBegin("HandleConnection", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProtocolServiceHandleConnectionArgs struct {
	Data string `thrift:"data,1" frugal:"1,default,string" json:"data"`
}

func NewProtocolServiceHandleConnectionArgs() *ProtocolServiceHandleConnectionArgs {
	return &ProtocolServiceHandleConnectionArgs{}
}

func (p *ProtocolServiceHandleConnectionArgs) InitDefault() {
	*p = ProtocolServiceHandleConnectionArgs{}
}

func (p *ProtocolServiceHandleConnectionArgs) GetData() (v string) {
	return p.Data
}
func (p *ProtocolServiceHandleConnectionArgs) SetData(val string) {
	p.Data = val
}

var fieldIDToName_ProtocolServiceHandleConnectionArgs = map[int16]string{
	1: "data",
}

func (p *ProtocolServiceHandleConnectionArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ProtocolServiceHandleConnectionArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ProtocolServiceHandleConnectionArgs) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Data = v
	}
	return nil
}

func (p *ProtocolServiceHandleConnectionArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("HandleConnection_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ProtocolServiceHandleConnectionArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("data", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Data); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *ProtocolServiceHandleConnectionArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ProtocolServiceHandleConnectionArgs(%+v)", *p)

}

func (p *ProtocolServiceHandleConnectionArgs) DeepEqual(ano *ProtocolServiceHandleConnectionArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Data) {
		return false
	}
	return true
}

func (p *ProtocolServiceHandleConnectionArgs) Field1DeepEqual(src string) bool {

	if strings.Compare(p.Data, src) != 0 {
		return false
	}
	return true
}

type ProtocolServiceHandleConnectionResult struct {
}

func NewProtocolServiceHandleConnectionResult() *ProtocolServiceHandleConnectionResult {
	return &ProtocolServiceHandleConnectionResult{}
}

func (p *ProtocolServiceHandleConnectionResult) InitDefault() {
	*p = ProtocolServiceHandleConnectionResult{}
}

var fieldIDToName_ProtocolServiceHandleConnectionResult = map[int16]string{}

func (p *ProtocolServiceHandleConnectionResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err = iprot.Skip(fieldTypeId); err != nil {
			goto SkipFieldTypeError
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
SkipFieldTypeError:
	return thrift.PrependError(fmt.Sprintf("%T skip field type %d error", p, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ProtocolServiceHandleConnectionResult) Write(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin("HandleConnection_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ProtocolServiceHandleConnectionResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ProtocolServiceHandleConnectionResult(%+v)", *p)

}

func (p *ProtocolServiceHandleConnectionResult) DeepEqual(ano *ProtocolServiceHandleConnectionResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	return true
}
