// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/diagnostic/v1/Diagnostic.proto

package prixa_diagnostic_v1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *BotConversationRequest) Validate() error {
	if this.Reply != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Reply); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Reply", err)
		}
	}
	return nil
}
func (this *ReplyData) Validate() error {
	for _, item := range this.Preconditions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Preconditions", err)
			}
		}
	}
	return nil
}
func (this *PreconditionsDataRequest) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *PreconditionsDataRequestProp) Validate() error {
	return nil
}
func (this *BotConversationResponse) Validate() error {
	if this.Result != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Result); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Result", err)
		}
	}
	for _, item := range this.LogEvents {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LogEvents", err)
			}
		}
	}
	return nil
}
func (this *ResultData) Validate() error {
	if this.Messages != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Messages); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Messages", err)
		}
	}
	if this.Actions != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Actions); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Actions", err)
		}
	}
	for _, item := range this.Preconditions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Preconditions", err)
			}
		}
	}
	return nil
}
func (this *MessagesData) Validate() error {
	return nil
}
func (this *ActionData) Validate() error {
	for _, item := range this.Value {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
			}
		}
	}
	if this.DiagnosisResult != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DiagnosisResult); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DiagnosisResult", err)
		}
	}
	return nil
}
func (this *ValueData) Validate() error {
	return nil
}
func (this *DiagnosisResultData) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	if this.UserDetails != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UserDetails); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UserDetails", err)
		}
	}
	for _, item := range this.Profiles {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Profiles", err)
			}
		}
	}
	for _, item := range this.Symptoms {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Symptoms", err)
			}
		}
	}
	for _, item := range this.Diseases {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Diseases", err)
			}
		}
	}
	return nil
}
func (this *UserDetailsData) Validate() error {
	return nil
}
func (this *ProfileData) Validate() error {
	return nil
}
func (this *SymptomDesc) Validate() error {
	return nil
}
func (this *PotentialDisease) Validate() error {
	if this.Triage != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Triage); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Triage", err)
		}
	}
	for _, item := range this.Labs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Labs", err)
			}
		}
	}
	return nil
}
func (this *TriageResult) Validate() error {
	return nil
}
func (this *LabInfo) Validate() error {
	return nil
}
func (this *PreconditionsData) Validate() error {
	return nil
}
func (this *LogEvents) Validate() error {
	return nil
}
func (this *CreatePrixaSessionRequest) Validate() error {
	if this.UserData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UserData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UserData", err)
		}
	}
	return nil
}
func (this *CreatePrixaSessionResponse) Validate() error {
	if this.UserData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UserData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UserData", err)
		}
	}
	return nil
}
func (this *UserData) Validate() error {
	if this.Insurance != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Insurance); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Insurance", err)
		}
	}
	return nil
}
func (this *InsuranceStatus) Validate() error {
	return nil
}
func (this *SendEmailRequest) Validate() error {
	return nil
}
func (this *SendEmailResponse) Validate() error {
	if this.Message == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Message", fmt.Errorf(`value '%v' must not be an empty string`, this.Message))
	}
	return nil
}
func (this *SendSurveyRequest) Validate() error {
	if this.ApplicationId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ApplicationId", fmt.Errorf(`value '%v' must not be an empty string`, this.ApplicationId))
	}
	if this.Feedback == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Feedback", fmt.Errorf(`value '%v' must not be an empty string`, this.Feedback))
	}
	if this.PartnerId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartnerId))
	}
	if this.SessionId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SessionId", fmt.Errorf(`value '%v' must not be an empty string`, this.SessionId))
	}
	if this.UserId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must not be an empty string`, this.UserId))
	}
	return nil
}
func (this *FeedbackContentResponse) Validate() error {
	if this.Question == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Question", fmt.Errorf(`value '%v' must not be an empty string`, this.Question))
	}
	if this.Instruction == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Instruction", fmt.Errorf(`value '%v' must not be an empty string`, this.Instruction))
	}
	return nil
}
func (this *SendFeedbackRequest) Validate() error {
	if this.UserId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must not be an empty string`, this.UserId))
	}
	if this.PartnerId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartnerId))
	}
	if this.ApplicationId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ApplicationId", fmt.Errorf(`value '%v' must not be an empty string`, this.ApplicationId))
	}
	if this.SessionId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SessionId", fmt.Errorf(`value '%v' must not be an empty string`, this.SessionId))
	}
	if this.SymptomId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SymptomId", fmt.Errorf(`value '%v' must not be an empty string`, this.SymptomId))
	}
	if this.Question == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Question", fmt.Errorf(`value '%v' must not be an empty string`, this.Question))
	}
	if this.Detail == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Detail", fmt.Errorf(`value '%v' must not be an empty string`, this.Detail))
	}
	return nil
}
func (this *SendFeedbackResponse) Validate() error {
	return nil
}
func (this *GetDiseaseArticleRequest) Validate() error {
	return nil
}
func (this *GetDiseaseArticleResponse) Validate() error {
	return nil
}
